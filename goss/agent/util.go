package main

import (
	"os"
	"os/exec"
)

func fileExist(p string) bool {
	_, err := os.Stat(p)
	if err != nil {
		return false
	}

	return true
}

func gossExecutable(config map[string]string) string {
	path, ok := config["goss"]
	if ok {
		if fileExist(path) {
			return path
		}

		return ""
	}

	path, err := exec.LookPath("goss")
	if err != nil {
		return ""
	}

	if fileExist(path) {
		return path
	}

	return ""
}
