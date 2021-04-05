// Package converter is for image transformation.
package converter

import (
	"os"
	"image"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
	"path/filepath"
	"io/ioutil"
	"strconv"
	"errors"
)

// ConvertExt converts image files to specified extension in the src directory.
func ConvertExt(src, from, to string) (int, error) {
	from = strings.ToLower(from)
	to = strings.ToLower(to)

	if err := validateArgs(from, to); err != nil {
		return 0, err
	}

	fileNames := make(chan string)
	go func() {
		walkDir(src, from, fileNames)
		close(fileNames)
	}()

	fileCount := 0
	uniqueCheck := make(map[string]int)
	for fn := range fileNames {
		file, err := os.Open(fn)
		if err != nil {
			return fileCount, err
		}
		defer file.Close()
		img, _, err := image.Decode(file)
		if err != nil {
			return fileCount, err
		}

		fileName := fileName(fn)
		if _, ok := uniqueCheck[fileName]; !ok {
			uniqueCheck[fileName] = 0
		} else {
			uniqueCheck[fileName]++
			fileName = fileName + "(" + strconv.Itoa(uniqueCheck[fileName]) + ")"
		}

		dstfile, err := os.Create(fmt.Sprintf("output/%s.%s", fileName, to))
		if err != nil {
			return fileCount, err
		}
		defer dstfile.Close()

		switch to {
		case "jpeg", "jpg":
			err = jpeg.Encode(dstfile, img, nil)
		case "png":
			err = png.Encode(dstfile, img)
		}
		if err != nil {
			return fileCount, err
		}

		_, err = io.Copy(dstfile, file)
		if err != nil {
			return fileCount, err
		}
		fileCount++
	}
	return fileCount, nil
}

func walkDir(dir, ext string, fileNames chan<- string) {
	ue := strings.ToUpper(ext)
	for _, entry := range dirents(dir) {
		if !strings.HasSuffix(entry.Name(), ext) &&
			!strings.HasSuffix(entry.Name(), ue) &&
			!entry.IsDir() {
			continue
		}
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, ext, fileNames)
		} else {
			fileNames <- filepath.Join(dir, entry.Name())
		}
	}
}

func dirents(dir string) ([]os.FileInfo) {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	return entries
}

func fileName(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func validateArgs(from, to string) error {
	ae := allowedExt{"jpg", "jpeg", "png"}

	if to == from {
		err := errors.New("converter: from and to should be different")
		return err
	}
	if !ae.contains(to) {
		return fmt.Errorf("converter: from should be %s, your to is %s", ae, to)
	}
	if !ae.contains(from) {
		return fmt.Errorf("converter: from should be %s, your to is %s", ae, from)
	}

	return nil
}

type allowedExt []string

func (ae allowedExt) contains(item string) bool {
	set := make(map[string]struct{}, len(ae))
	for _, s := range ae {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
