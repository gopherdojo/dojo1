package converter

import (
	"fmt"
	"os"
	"path/filepath"
)

func convert(path string, converter *Converter) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}

	return filepath.Walk(path, walker(converter))
}

func walker(c *Converter) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer file.Close()

		img, err := c.Decode(file)
		if err != nil {
			return nil
		}

		newFile, err := os.Create(c.BuildNewFilePath(path))
		if err != nil {
			return err
		}
		defer newFile.Close()

		err = c.Encode(newFile, img)
		if err == nil {
			fmt.Printf("converted %s to %s\n", path, newFile.Name())
		}
		return err
	}
}
