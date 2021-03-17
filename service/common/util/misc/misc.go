package misc

import (
	"os"
)

// var storageConfig = config.GetStorageConfig()

func IsPathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}

	return false
}

func IsPathDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}

// func ZipFiles(sources []string, target string) error {
// 	zipfile, err := os.Create(target)
// 	if err != nil {
// 		return err
// 	}
// 	defer zipfile.Close()

// 	// Create a new zip archive.
// 	archive := zip.NewWriter(zipfile)
// 	defer archive.Close()

// 	rootDir := storageConfig.StorageRootDir

// 	for _, path := range sources {
// 		fullpath := rootDir + "/" + path
// 		isDir, err := IsPathDir(fullpath)
// 		if err != nil {
// 			return err
// 		}

// 		if !isDir {
// 			continue
// 		}

// 		baseDir := filepath.Base(fullpath)
// 		filepath.Walk(fullpath, func(path string, info os.FileInfo, err error) error {
// 			header, err := zip.FileInfoHeader(info)
// 			if err != nil {
// 				return err
// 			}

// 			if len(baseDir) > 0 {
// 				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, fullpath))
// 			}

// 			if info.IsDir() {
// 				header.Name += "/"
// 			} else {
// 				header.Method = zip.Deflate
// 			}

// 			writer, err := archive.CreateHeader(header)
// 			if err != nil {
// 				return err
// 			}

// 			if info.IsDir() {
// 				return nil
// 			}

// 			file, err := os.Open(path)
// 			if err != nil {
// 				return err
// 			}
// 			defer file.Close()
// 			_, err = io.Copy(writer, file)
// 			return err
// 		})
// 	}

// 	return nil
// }
