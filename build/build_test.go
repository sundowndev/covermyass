package build

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestBuild(t *testing.T) {
	t.Run("version and commit default values", func(t *testing.T) {
		assert.Equal(t, "dev", version)
		assert.Equal(t, "dev", commit)
		assert.Equal(t, false, IsRelease())
		assert.Equal(t, "dev-dev", Name())
	})

	t.Run("version and commit default values", func(t *testing.T) {
		version = "v2.4.4"
		commit = "0ba854f"
		assert.Equal(t, true, IsRelease())
		assert.Equal(t, "v2.4.4-0ba854f", Name())

		// Reset values
		version = "dev"
		commit = "dev"
	})

	t.Run("version and commit with Go version", func(t *testing.T) {
		assert.Equal(t, false, IsRelease())
		assert.Equal(t, fmt.Sprintf("dev-dev (%s)", runtime.Version()), String())
	})
}
