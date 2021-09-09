package path

import (
	"os"
	"path/filepath"
)

func getRootPath() string {
	path, _ := os.Getwd()
	return filepath.Join(path, "../..")
}

func GetFilePath(path string) string {
	return filepath.Join(getRootPath(), path)
}
