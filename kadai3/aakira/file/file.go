package file

import (
	"os"
	"fmt"
	"bufio"
)

func ReadFile(path string, length int) []string {
	var fp *os.File
	var err error

	fp, err = os.Open(path)

	if err != nil {
		fmt.Printf("File not found : %s\n", path)
		return nil
	}
	defer fp.Close()

	var lines []string
	scanner := bufio.NewScanner(fp)
	for i := 0; scanner.Scan() && i < length; i++ {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	return lines
}
