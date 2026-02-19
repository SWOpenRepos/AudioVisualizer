//go:build linux
// +build linux

package window

import (
	"SWOpen/AudioVisualizer/internal/cmd"
)

func openURL(url string) error {
	args := []string{url}

	return cmd.Run("xdg-open", nil, args...)
}
