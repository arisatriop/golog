package golog

import (
	"testing"
)

func TestLogError(t *testing.T) {
	LogError("example activity", "example file location", "example error message", "example user")
}

func TestLogStart(t *testing.T) {
	t.Run("", func(t *testing.T) {
		LogStart("example activity", "example users")
	})

	t.Run("", func(t *testing.T) {
		LogStart("example activity")
	})

	t.Run("", func(t *testing.T) {
		LogStart()
	})
}

func TestLogEnd(t *testing.T) {
	LogEnd("example activity", "example users")
}
