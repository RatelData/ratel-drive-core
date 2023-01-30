package util

import (
	"os"

	"go.uber.org/zap"
)

func IsPathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}

	return false
}

func IsPathDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fi.IsDir()
}

func CheckCreateDataDirectory() {
	logger := GetLogger()

	rootDir := GetStorageConfig().StorageRootDir
	if err := os.MkdirAll(rootDir, os.ModePerm); err != nil {
		logger.Error("Create data directory failed!",
			zap.String("error", "Please check if you have the permission!"),
		)
		return
	}
}
