package main

import (
	"fmt"
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

// ensureBinaries ensures all necessary binaries are available
func ensureBinaries() error {
	errorStr := ""

	for _, binary := range []string{"bash", "file", "find"} {
		_, err := exec.Command("where", "/F", binary).Output()
		if err != nil {
			errorStr += binary + ".exe "
		}
	}

	if errorStr != "" {
		return fmt.Errorf(MESSAGES["WindowsToolingNotFound"] + errorStr)
	}

	return nil
}

func platformChecks() error {
	if binaryErr := ensureBinaries(); binaryErr != nil {
		return binaryErr
	}

	return nil
}
