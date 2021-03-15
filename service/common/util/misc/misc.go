package misc

import "os"

func IsPathExists(path string) bool {
	if _, err := os.Stat("path"); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}

	return false
}
