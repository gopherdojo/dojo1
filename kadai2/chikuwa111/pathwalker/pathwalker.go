// Package pathwalker provides Find function.
package pathwalker

import (
	"os"
	"path/filepath"
	"strings"
)

// Find search files that have the extension in path.
// When files are found, this func passes path handler.
func Find(path string, extension string, handler func(string) error) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.ToLower(filepath.Ext(path)) != "."+extension {
			return nil
		}
		handlerErr := handler(path)
		return handlerErr
	})
	return err
}
