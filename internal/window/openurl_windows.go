//go:build windows
// +build windows

package window

import (
	"SWOpen/AudioVisualizer/internal/cmd"
)

func openURL(url string) error {
	args := []string{"/c", "start", url}

	return cmd.Run("cmd", nil, args...)
}
