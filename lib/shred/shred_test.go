package shred

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/sundowndev/covermyass/v2/mocks"
	"runtime"
	"testing"
)

func TestShredder_Write(t *testing.T) {
	cases := []struct {
		name    string
		options ShredderOptions
		input   string
		wantErr map[string]error
	}{
		{
			name:  "test with non-existing file",
			input: "testdata/fake.log",
			wantErr: map[string]error{
				"linux":   errors.New("file stat failed: stat testdata/fake.log: no such file or directory"),
				"windows": errors.New("file stat failed: CreateFile testdata/fake.log: The system cannot find the file specified."),
			},
		},
		{
			name:  "test with non-file path",
			input: "testdata/",
			wantErr: map[string]error{
				"linux":   errors.New("file opening failed: open testdata/: is a directory"),
				"windows": errors.New("file opening failed: open testdata/: is a directory"),
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(&tt.options)

			err := s.Write(tt.input)
			if tt.wantErr[runtime.GOOS] == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr[runtime.GOOS].Error())
			}
		})
	}
}

func TestShredder_shred(t *testing.T) {
	cases := []struct {
		name      string
		options   ShredderOptions
		mocks     func(*mocks.FileInfo, *mocks.File)
		wantError error
	}{
		{
			name: "test writing empty file",
			options: ShredderOptions{
				Zero:       false,
				Iterations: 3,
			},
			mocks: func(fakeFileInfo *mocks.FileInfo, fakeFile *mocks.File) {
				fakeFileInfo.On("Size").Return(int64(0)).Times(1)
			},
		},
		{
			name: "test writing a 64 bytes file",
			options: ShredderOptions{
				Zero:       false,
				Iterations: 3,
			},
			mocks: func(fakeFileInfo *mocks.FileInfo, fakeFile *mocks.File) {
				fakeFileInfo.On("Size").Return(int64(64)).Times(4)

				fakeFile.On("Seek", int64(0), 0).Return(int64(0), nil).Times(3)
				fakeFile.On("Sync").Return(nil).Times(3)
				fakeFile.On("Write", mock.MatchedBy(func(b []byte) bool {
					return len(b) != 0
				})).Return(0, nil)
			},
		},
		{
			name: "test writing a 2Mb file with 10 iterations",
			options: ShredderOptions{
				Zero:       false,
				Iterations: 10,
			},
			mocks: func(fakeFileInfo *mocks.FileInfo, fakeFile *mocks.File) {
				fakeFileInfo.On("Size").Return(int64(2000000)).Times(11)

				fakeFile.On("Seek", int64(0), 0).Return(int64(0), nil).Times(10)
				fakeFile.On("Sync").Return(nil).Times(10)
				fakeFile.On("Write", mock.MatchedBy(func(b []byte) bool {
					return len(b) != 0
				})).Return(0, nil)
			},
		},
		{
			name: "test writing a 2Kb file with error",
			options: ShredderOptions{
				Zero:       false,
				Iterations: 3,
			},
			mocks: func(fakeFileInfo *mocks.FileInfo, fakeFile *mocks.File) {
				fakeFileInfo.On("Size").Return(int64(2000)).Times(2)

				fakeFile.On("Seek", int64(0), 0).Return(int64(0), nil).Times(1)
				fakeFile.On("Write", mock.MatchedBy(func(b []byte) bool {
					return len(b) != 0
				})).Return(0, errors.New("dummy error"))
			},
			wantError: errors.New("dummy error"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(&tt.options)

			fakeFileInfo := &mocks.FileInfo{}
			fakeFile := &mocks.File{}
			tt.mocks(fakeFileInfo, fakeFile)

			err := s.shred(fakeFileInfo, fakeFile)
			if tt.wantError == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError.Error())
			}

			fakeFileInfo.AssertExpectations(t)
			fakeFile.AssertExpectations(t)
		})
	}
}
