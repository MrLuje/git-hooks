package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func isBashExecutable(file string) bool {
	return true

	mime, err := getFileMimeType(file)
	if err != nil {
		return false
	}
	return strings.Contains(mime, "text/x-shellscript")
}

func getFileMimeType(path string) (string, error) {
	b, err := exec.Command("file", []string{path, "--mime-type"}...).Output()
	if err == nil {
		fmt.Printf("%v checking mime type %v\n", path, string(b))
		return string(b), nil
	}
	return "", err
}

func isExecutable(info os.FileInfo, folder string) bool {
	fileName := info.Name()
	b, err := exec.Command("find", []string{filepath.Join(folder, fileName), "-maxdepth", "1", "-executable", "-type", "f"}...).Output()
	if err != nil {
		return false
	}

	// find _path_ -executable returns an empty string if not en executable
	return string(b) != ""
}

// func isExecutablea(info os.FileInfo, folder string) bool {
// 	name := info.Name()
// 	if strings.Contains(name, ".") {
// 		fmt.Printf("%v contains a . OK\n", name)
// 		return true
// 	}
// 	if isBashExecutable(filepath.Join(folder, name)) {
// 		return true
// 	}

// 	fmt.Printf("%v NO IDEA SO OK\n", name)

// 	return true // no way to properly check this...
// }
