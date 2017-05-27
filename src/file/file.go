package file

import (
	"os"
	"path/filepath"
	"strings"
)

// get all files in the path
func GetFilelist(path string) ([]string, *error) {
	fileV := []string{}
	var e *error = nil
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			*e = err
			return nil
		}
		if f.IsDir() {
			return nil
		}
		fileV = append(fileV, path)
		return nil
	})
	return fileV, e
}

// convert path to useful
func FixPath(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}
