package golog

import (
	"testing"
)

func TestLogError(t *testing.T) {
	LogError("example activity", "example file location", "example error message", "example user")
}

func TestLogStart(t *testing.T) {
	LogStart("exmaple activity", "example users")
}

func TestLogEnd(t *testing.T) {
	LogEnd("exmaple activity", "example users")
}
