package converter

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Converter has image encode/decode function, and conversion destination format.
type Converter struct {
	Decoder
	Encoder

	ext string
}

// BuildNewFilePath creates conversion destination file path.
func (c *Converter) BuildNewFilePath(path string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + "." + c.ext
}

// Convert converts image files under the path recursively.
func (c *Converter) Convert(path string) error {
	return convert(path, c)
}

// NewConverter creates Converter.
func NewConverter(from, to string) (*Converter, error) {
	d, err := NewDecoder(from)
	if err != nil {
		return nil, err
	}

	e, err := NewEncoder(to)
	if err != nil {
		return nil, err
	}

	if d.ext() == e.ext() {
		return nil, fmt.Errorf("invalid same format: from=%s, to=%s", from, to)
	}

	return &Converter{d, e, to}, nil
}
