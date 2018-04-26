package pathwalker_test

import (
	"pathwalker"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFind(t *testing.T) {
	cases := []struct {
		name      string
		path      string
		extension string
		expected  []string
	}{
		{name: "txtInFiles1", path: "testdata/files1", extension: "txt", expected: []string{"testdata/files1/text1.txt", "testdata/files1/text2.txt"}},
		{name: "logInFiles1", path: "testdata/files1", extension: "log", expected: []string{"testdata/files1/log1.log", "testdata/files1/log2.log"}},
		{name: "recursive", path: "testdata", extension: "txt", expected: []string{"testdata/files1/text1.txt", "testdata/files1/text2.txt", "testdata/files2/text1.txt", "testdata/files2/text2.txt"}},
		{name: "filePath", path: "testdata/files1/text1.txt", extension: "txt", expected: []string{"testdata/files1/text1.txt"}},
		{name: "noFiles", path: "testdata/no-files", extension: "txt", expected: []string{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := []string{}
			err := pathwalker.Find(c.path, c.extension, func(path string) error {
				actual = append(actual, path)
				return nil
			})
			if err != nil {
				t.Errorf("Got error (args path: %s, extension: %s)", c.path, c.extension)
			}
			if !cmp.Equal(actual, c.expected) {
				t.Errorf("Expected: %v, got %v (args path: %s, extension: %s)", c.expected, actual, c.path, c.extension)
			}
		})
	}
}
