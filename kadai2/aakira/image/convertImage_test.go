package image

import (
	"testing"
)

func TestToImage(t *testing.T) {
	t.Helper()

	extension := "jpg"
	t.Run("Expect " + extension , func(t *testing.T) {
		files := []struct {
			path  string
		}{
			{"hoge.jpg"},
			{"foooooo.png.jpg"},
		}

		for _, test := range files {
			actual, err := ToImageFile(test.path)

			// pass: file not found
			if err != nil {
				return
			}

			if v, ok := actual.(*JpgImage); !ok {
				t.Fatalf("actual: %v\n expected: %v\n", v, extension)
			}
		}
	})

	extension = "png"
	t.Run("Expect " + extension , func(t *testing.T) {
		files := []struct {
			path  string
		}{
			{"hoge.png"},
			{"foooooo.jpg.png"},
		}

		for _, test := range files {
			actual, err := ToImageFile(test.path)

			// pass: file not found
			if err != nil {
				return
			}

			if v, ok := actual.(*PngImage); !ok {
				t.Fatalf("actual: %v\n expected: %v\n", v, extension)
			}
		}
	})

	t.Run("Expect nil", func(t *testing.T) {
		files := []struct {
			path  string
		}{
			{"hoge"},
			{"foo.webp"},
			{""},
		}

		for _, test := range files {
			actual, err := ToImageFile(test.path)

			if err == nil {
				t.Fatalf("actual: %v\n expected: nil\n", actual)
			}
		}
	})
}
