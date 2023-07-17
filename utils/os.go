package AnonUtils

import "os"

func IsDir(path string) bool {
	drive, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	if !drive.IsDir() {
		return false
	}
	return true
}
