package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	testcases := []struct {
		name  string
		rules []string
		paths map[string]bool
	}{
		{
			name: "test with simple patterns",
			rules: []string{
				"/var/*",
				"/foo/**/*",
				"/bar/foo/1.log",
			},
			paths: map[string]bool{
				"/var/log/lastlog":       false,
				"/var/test.log":          true,
				"/var/fakefile":          true,
				"/db/logfile.log":        false,
				"/foo/bar/1/logfile.log": true,
				"/bar/foo/1.log":         true,
				"/var/**/*.log":          false,
			},
		},
		{
			name: "test patterns against patterns",
			rules: []string{
				"/var/*",
				"/var/*/foo/**/*",
				"/foo/**/*",
				"/bar/foo/1.log",
			},
			paths: map[string]bool{
				"/var/*.log":        true,
				"/var/**/*":         false,
				"/foo/*.log":        true,
				"/bar/foo/**/*.log": false,
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			f := NewEngine()
			err := f.AddRule(tt.rules...)
			assert.NoError(t, err)

			for path, shouldMatch := range tt.paths {
				assert.Equal(t, shouldMatch, f.Match(path))
			}
		})
	}
}
