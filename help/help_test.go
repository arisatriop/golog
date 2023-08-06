package help

import (
	"fmt"
	"testing"
)

func TestGetTime(t *testing.T) {
	t.Run("", func(t *testing.T) {
		now := GetTime()
		fmt.Println(now)
	})
	t.Run("", func(t *testing.T) {
		now := GetTime("UTC")
		fmt.Println(now)
	})
	t.Run("", func(t *testing.T) {
		now := GetTime("UTCs")
		fmt.Println(now)
	})
}
