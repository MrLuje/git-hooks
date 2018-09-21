package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func isExecutable(info os.FileInfo, folder string) bool {
	name := info.Name()
	if strings.Contains(name, ".") {
		fmt.Printf("%v contains a . OK\n", name)
		return true
	}
	if b, err := exec.Command("file", []string{filepath.Join(folder, name), "--mime-type"}...).Output(); err == nil {
		fmt.Printf("%v checking mime type %v\n", filepath.Join(folder, name), string(b))
		return strings.Contains(string(b), "text/x-shellscript")
	}

	fmt.Printf("%v NO IDEA SO OK\n", name)

	return true // no way to properly check this...
}
