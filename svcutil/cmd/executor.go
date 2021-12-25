package cmd

import (
	"context"
	"os/exec"
)

type terminal struct {
	Cmd string
}

type Terminal interface {
	ExecuteBash(context.Context, ...string) (string, error)
}

func NewTerminal(command string) Terminal {
	return &terminal{
		Cmd: command,
	}
}

func (sh *terminal) ExecuteBash(ctx context.Context, args ...string) (output string, err error) {
	cmd := exec.Command(sh.Cmd, args...)

	stdout, err := cmd.Output()
	if err != nil {
		return output, err
	}

	output = string(stdout)

	return output, err
}
