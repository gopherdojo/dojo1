package file

import (
	"path/filepath"
	"os"
	"fmt"
	"errors"
)

/*
find files in specified directory

throw error if file not found.
 */
func FindFiles(path string, fileExtension string) ([]string, error) {
	var imagePaths []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == fileExtension {
			imagePaths = append(imagePaths, path)
		}
		return err
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if len(imagePaths) == 0 {
		return nil, errors.New("file not found")
	}
	return imagePaths, nil
}
