package analysis

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/sundowndev/covermyass/v2/mocks"
	"os"
	"testing"
)

func TestAnalysis_Write(t *testing.T) {
	fakeCheck := &mocks.Check{}

	testcases := []struct {
		name     string
		filename string
		results  []Result
	}{
		{
			name:     "test valid analysis",
			filename: "testdata/valid_analysis.txt",
			results: []Result{
				{
					Path:     "/var/log/dummy.log",
					Size:     int64(32),
					Mode:     os.ModePerm,
					ReadOnly: false,
					Check:    fakeCheck,
				},
				{
					Path:     "/var/log/dummy2.log",
					Size:     int64(8),
					Mode:     600,
					ReadOnly: true,
					Check:    fakeCheck,
				},
			},
		},
		{
			name:     "test empty analysis",
			filename: "testdata/empty_analysis.txt",
			results:  []Result{},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			testdata, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Fatal(err)
			}

			a := NewAnalysis()
			for _, r := range tt.results {
				a.AddResult(r)
			}
			assert.Len(t, a.Results(), len(tt.results))

			buffer := &bytes.Buffer{}
			a.Write(buffer)

			assert.Equal(t, string(testdata), buffer.String())
		})
	}

	fakeCheck.AssertExpectations(t)
}
