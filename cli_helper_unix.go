// +build !windows

package main

import (
	"os/exec"
)

func prepareCmd(hook string, args []string) *exec.Cmd {
	return exec.Command(hook, args...)
}
