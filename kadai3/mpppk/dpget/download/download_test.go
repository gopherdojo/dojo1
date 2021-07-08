package download

import (
	"testing"
)

func TestGenerateRangeHeaders(t *testing.T) {
	cases := []struct {
		contentLength int
		splitNum      int
		expected      []string
	}{
		{
			contentLength: 1,
			splitNum:      1,
			expected:      []string{"bytes=0-0"},
		},
		{
			contentLength: 2,
			splitNum:      1,
			expected:      []string{"bytes=0-1"},
		},
		{
			contentLength: 2,
			splitNum:      2,
			expected:      []string{"bytes=0-0", "bytes=1-1"},
		},
		{
			contentLength: 3,
			splitNum:      2,
			expected:      []string{"bytes=0-0", "bytes=1-2"},
		},
	}

	for _, c := range cases {
		headers := generateRangeHeaders(c.contentLength, c.splitNum)

		if len(headers) != len(c.expected) {
			t.Fatalf("generateRangeHeaders is expected to return %d elements when contentLength is %d, but actually returns %d elements",
				len(c.expected), c.contentLength, len(headers))
		}

		for i, header := range headers {
			if header != c.expected[i] {
				t.Fatalf("generateRangeHeaders is expected to return %q as an element with index %d if contentLength is %d, but acutually returns %q.",
					c.expected, i, c.contentLength, header)
			}
		}
	}
}
