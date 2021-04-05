package main

import (
	"os"
	"github.com/prometheus/common/log"
	"bufio"
	"unicode"
	"fmt"
)

// a tool for generating a words file from the user dict
func main() {
	path := "/usr/share/dict/words"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal()
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) > 10 && !isFirstUpper(t) {
			lines = append(lines, t)
		}
	}

	writer, err := os.Create("words.txt")
	defer writer.Close()
	if err != nil {
		log.Fatal()
	}

	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
}

func isFirstUpper(v string) bool {
	if v == "" {
		return false;
	}
	r := rune(v[0])
	return unicode.IsUpper(r)
}
