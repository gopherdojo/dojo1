package file

import (
	"testing"
)

func TestFindFiles(t *testing.T) {
	t.Helper()

	t.Run("Directory not found ", func(t *testing.T) {
		actual, err := FindFiles("./testdata", ".jpg")
		if err == nil {
			t.Fatalf("actual: %v\n expected: directory not found error", actual)
		}
	})

	t.Run("File not found ", func(t *testing.T) {
		files := []struct {
			extension string
			hasError  bool
		}{
			{".jpg", false},
			{".webp", true},
			{".go", true},
		}
		for _, test := range files {
			actual, err := FindFiles("../testdata", test.extension)

			if (test.hasError && err == nil) || (!test.hasError && err != nil) {
				t.Fatalf("actual: %v\n", actual)
			}
		}
	})
}
