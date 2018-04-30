package typing

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"time"
)

// MakeWords create word list from text file.
func MakeWords(wordListPath string) ([]string, error) {
	file, err := os.Open(wordListPath)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 4096)
	var words []string
	for {
		line, _, err := reader.ReadLine()
		words = append(words, string(line))
		if err == io.EOF {
			break
		} else if err != nil {
			return []string{}, err
		}
	}
	shuffle(words)
	return words, nil
}

func shuffle(xs []string) {
	// ref: http://marcelom.github.io/2013/06/07/goshuffle.html
	rand.Seed(time.Now().UnixNano())
	for i := range xs {
		j := rand.Intn(i + 1)
		xs[i], xs[j] = xs[j], xs[i]
	}
}
