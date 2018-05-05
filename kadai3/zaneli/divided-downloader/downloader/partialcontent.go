package downloader

import "fmt"

// PartialContent has Range header's from and to byte position.
type PartialContent struct {
	from  int
	to    int
	index int
	path  string
}

func (p PartialContent) filePath() string {
	return fmt.Sprintf("%s.%d", p.path, p.index)
}
