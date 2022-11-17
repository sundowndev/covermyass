package shred

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/sundowndev/covermyass/v2/mocks"
	"testing"
)

func TestShredder_Write(t *testing.T) {
	cases := []struct {
		name      string
		options   ShredderOptions
		input     string
		wantError error
	}{
		{
			name:      "test with non-existing file",
			input:     "testdata/fake.log",
			wantError: errors.New("shredding failed: stat testdata/fake.log: no such file or directory"),
		},
		{
			name:      "test with non-file path",
			input:     "testdata/",
			wantError: errors.New("shredding failed: open testdata/: is a directory"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(&tt.options)

			err := s.Write(tt.input)
			if tt.wantError == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError.Error())
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
				Unlink:     false,
			},
			mocks: func(fakeFileInfo *mocks.FileInfo, fakeFile *mocks.File) {
				fakeFileInfo.On("Size").Return(int64(0)).Times(1)
				fakeFile.On("Close").Return(nil).Times(1)
			},
		},
		{
			name: "test writing a 64 bytes file",
			options: ShredderOptions{
				Zero:       false,
				Iterations: 3,
				Unlink:     false,
			},
			mocks: func(fakeFileInfo *mocks.FileInfo, fakeFile *mocks.File) {
				fakeFileInfo.On("Size").Return(int64(64)).Times(4)

				fakeFile.On("Seek", int64(0), 0).Return(int64(0), nil).Times(3)
				fakeFile.On("Close").Return(nil).Times(1)
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
				Unlink:     false,
			},
			mocks: func(fakeFileInfo *mocks.FileInfo, fakeFile *mocks.File) {
				fakeFileInfo.On("Size").Return(int64(2000000)).Times(11)

				fakeFile.On("Seek", int64(0), 0).Return(int64(0), nil).Times(10)
				fakeFile.On("Close").Return(nil).Times(1)
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
				Unlink:     false,
			},
			mocks: func(fakeFileInfo *mocks.FileInfo, fakeFile *mocks.File) {
				fakeFileInfo.On("Size").Return(int64(2000)).Times(2)

				fakeFile.On("Seek", int64(0), 0).Return(int64(0), nil).Times(1)
				fakeFile.On("Close").Return(nil).Times(1)
				fakeFile.On("Write", mock.MatchedBy(func(b []byte) bool {
					return len(b) != 0
				})).Return(0, errors.New("dummy error"))
			},
			wantError: errors.New("dummy error"),
		},
		{
			name: "test writing a 2Kb file with zero option",
			options: ShredderOptions{
				Zero:       true,
				Iterations: 5,
				Unlink:     false,
			},
			mocks: func(fakeFileInfo *mocks.FileInfo, fakeFile *mocks.File) {
				fakeFileInfo.On("Size").Return(int64(2000)).Times(6)

				fakeFile.On("Close").Return(nil).Times(1)
				fakeFile.On("Seek", int64(0), 0).Return(int64(0), nil).Times(5)
				fakeFile.On("Sync").Return(nil).Times(5)
				fakeFile.On("Write", mock.MatchedBy(func(b []byte) bool {
					return len(b) > 0
				})).Return(0, nil)
				fakeFile.On("Write", mock.MatchedBy(func(b []byte) bool {
					return len(b) == 0
				})).Return(0, nil).Once()
			},
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
