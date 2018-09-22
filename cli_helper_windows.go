package main

import (
	"os/exec"
	"strings"
)

func toPOSIX(hook string) string {
	return "/" + strings.Replace(strings.Replace(hook, ":\\", "/", -1), "\\", "/", -1)
}

func prepareCmd(hook string, args []string) *exec.Cmd {
	if isBashExecutable(hook) {
		return asBashCmd(hook, args)
	}

	return exec.Command(hook, args...)
}

// asBashCmd execute the hook through bash -c
func asBashCmd(hook string, args []string) *exec.Cmd {
	_hook := toPOSIX(hook)
	_args := append([]string{"-c", _hook}, args...)

	return exec.Command("bash", _args...)
}
