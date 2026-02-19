package cmd

import (
	"bytes"
	"os/exec"
)

func Run(command string, buf *bytes.Buffer, args ...string) error {
	cmd := exec.Command(command, args...)
	setProcAttr(cmd)

	if buf != nil {
		cmd.Stdout = buf
	}

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func Load(command string, args []string) (cmd *exec.Cmd) {
	cmd = exec.Command(command, args...)
	setProcAttr(cmd)

	return cmd
}
