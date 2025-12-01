package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Uses xclip on linux
func CopyToClipboard(text string) error {
	command := exec.Command("wl-copy")
	command.Stdin = bytes.NewReader([]byte(text))

	if err := command.Start(); err != nil {
		return fmt.Errorf("error starting xclip command: %w", err)
	}

	if err := command.Wait(); err != nil {
		return fmt.Errorf("error running xclip: %w", err)
	}

	return nil
}
