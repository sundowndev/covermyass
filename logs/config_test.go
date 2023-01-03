package logs

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfig_GetConfig(t *testing.T) {
	t.Run("test default behavior", func(t *testing.T) {
		Init()
		config := getConfig()
		assert.Equal(t, logrus.DebugLevel, config.Level)
	})

	t.Run("test custom log level", func(t *testing.T) {
		_ = os.Setenv("LOG_LEVEL", "warn")
		defer os.Clearenv()

		Init()
		config := getConfig()
		assert.Equal(t, logrus.WarnLevel, config.Level)
	})
}
