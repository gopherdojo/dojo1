package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	_ "golang.org/x/image/webp"
	"os"
	"path/filepath"
	"strings"
)

// formatInspecter inspects supported image format.
type formatInspecter interface {
	Inspect(string) bool
}

// Formats is the list of registered image formats.
type Formats []string

// Inspect returns true value when image format is supported.
func (f *Formats) Inspect(file string) bool {
	for _, format := range *f {
		if format == strings.TrimLeft(filepath.Ext(file), ".") {
			return true
		}
	}
	return false
}

// SourceFormats is the list of supported source formats.
var SourceFormats = Formats{"png", "jpg", "gif", "webp"}

// DestFormats is the list of supported destination formats.
var DestFormats = Formats{"png", "jpg", "gif"}

// target is the pair of conversion files
type target struct {
	src, dest string
}

// GetSrc returns property of src
func (t *target) GetSrc() string {
	return t.src
}

// GetDest returns property of dest
func (t *target) GetDest() string {
	return t.dest
}

// convert executes image conversion on target
func (t *target) convert() error {
	if !SourceFormats.Inspect(t.src) {
		return fmt.Errorf("src:%v is not supported", t.src)
	}
	if !DestFormats.Inspect(t.dest) {
		return fmt.Errorf("dest:%s is not supported", t.dest)
	}

	file, err := os.Open(t.src)
	if err != nil {
		return err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	w, err := os.Create(t.dest)
	if err != nil {
		return err
	}
	defer w.Close()

	switch filepath.Ext(t.dest) {
	case ".png":
		return png.Encode(w, img)
	case ".jpg":
		return jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
	case ".gif":
		return gif.Encode(w, img, &gif.Options{NumColors: 256})
	}
	return fmt.Errorf("Unknown error: %v", t.dest) // FIXME: panic でもいいのかも
}

// Convert executes image conversion a source file to the destination file.
func Convert(src, dest string) error {
	t := &target{src, dest}
	return t.convert()
}

// RecursiveConverter converts target images recursively.
type RecursiveConverter struct {
	// in is input directory.
	in string
	// out is output directory.
	out string
	// srcFormat is image format before conversion.
	srcFormat string
	// to is image format after conversion.
	destFormat string
	// targets
	targets []*target
}

// NewRecursiveConverter allocates a new RecursiveConverter struct and detect error.
func NewRecursiveConverter(in, out, srcFormat, destFormat string) (*RecursiveConverter, error) {
	rc := &RecursiveConverter{}

	for _, dir := range []string{in, out} {
		stat, err := os.Stat(dir)
		if err != nil {
			return rc, err
		}
		if !stat.IsDir() {
			return rc, fmt.Errorf("%s is not directory", dir)
		}
	}

	if srcFormat == destFormat {
		return rc, fmt.Errorf("same formats are specified")
	}

	rc.in = in
	rc.out = out
	rc.srcFormat = srcFormat
	rc.destFormat = destFormat
	rc.buildTargets()

	return rc, nil
}

// GetTargets returns property of targets.
func (rc *RecursiveConverter) GetTargets() []*target {
	return rc.targets
}

// Convert executes image conversion for target files.
func (rc *RecursiveConverter) Convert() error {
	for _, t := range rc.targets {
		if err := t.convert(); err != nil {
			return err
		}
	}
	return nil
}

// buildTargets is store the pair of target to targets.
func (rc *RecursiveConverter) buildTargets() error {
	err := filepath.Walk(rc.in, func(src string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if rc.srcFormat != strings.TrimLeft(filepath.Ext(src), ".") {
			return nil
		}
		t := &target{src, rc.buildDest(src)}
		rc.targets = append(rc.targets, t)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// buildDestPath creates the destination file path.
func (rc *RecursiveConverter) buildDest(src string) string {
	in := filepath.Base(rc.in)
	srcAbs, _ := filepath.Abs(src)
	srcElems := strings.Split(srcAbs, string(os.PathSeparator))

	destElems := []string{rc.out}
	for i, elem := range srcElems {
		if in == elem {
			destElems = append(destElems, srcElems[i+1:]...)
			continue
		}
	}
	basename := filepath.Base(src)
	destElems[len(destElems)-1] = strings.TrimSuffix(basename, filepath.Ext(basename)) + "." + rc.destFormat
	dest := filepath.Join(destElems...)

	destDir := filepath.Dir(dest)
	if _, err := os.Stat(destDir); err != nil {
		os.MkdirAll(destDir, os.ModePerm)
	}
	return dest
}
