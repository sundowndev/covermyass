package analysis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_byteCountSI(t *testing.T) {
	testcases := map[int64]string{
		0:                   "0 B",
		1:                   "1 B",
		10000:               "10.0 kB",
		22736526:            "22.7 MB",
		2123652683:          "2.1 GB",
		2123652683000:       "2.1 TB",
		2123652683000000:    "2.1 PB",
		2123652683000000000: "2.1 EB",
	}

	for v, expected := range testcases {
		t.Run(fmt.Sprintf("test with input %d", v), func(t *testing.T) {
			assert.Equal(t, expected, byteCountSI(v))
		})
	}
}
