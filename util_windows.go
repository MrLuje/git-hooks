package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// isBashExecutable tells if this file can be executed with bash
func isBashExecutable(file string) bool {
	mime, err := getFileMimeType(file)
	if err != nil {
		return false
	}
	return strings.Contains(mime, "text/x-shellscript")
}

// getFileMimeType gets mime-type using file.exe
func getFileMimeType(path string) (string, error) {
	b, err := exec.Command("file", []string{path, "--mime-type"}...).Output()
	if err != nil {
		logger.Infoln(fmt.Sprintf("Can't get mime-type of %v: %s", path, err))
		return "", err
	}

	return string(b), nil
}

// getFindLocation searches for find.exe with 'git' in its path
func getFindLocation() (string, error) {
	b, err := exec.Command("where", "find").Output()
	if err != nil {
		logger.Infoln("Can't get location of find.exe: " + err.Error())
		return "", err
	}

	locations := string(b)
	for _, location := range strings.Split(locations, "\r\n") {
		if strings.Contains(strings.ToLower(location), "git") {
			return location, nil
		}
	}

	return "", fmt.Errorf("Not found")
}

func isExecutable(info os.FileInfo, folder string) bool {
	fileName := info.Name()
	findPath, err := getFindLocation()

	if err != nil {
		return false
	}

	fullpathToFile := filepath.Join(folder, fileName)
	b, err := exec.Command(findPath, []string{fullpathToFile, "-maxdepth", "1", "-executable", "-type", "f"}...).Output()
	if err != nil {
		logger.Infoln(fmt.Sprintf("Can't ensure %v is an executable: %s", fullpathToFile, err))
		return false
	}

	// find _path_ -executable returns an empty string if not en executable
	return string(b) != ""
}
