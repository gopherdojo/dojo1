package httpHeader

import (
	"strconv"
	"net/http"
)

func GetLength(url string) (int, error) {
	response, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	maps := response.Header
	length, err := strconv.Atoi(maps["Content-Length"][0])
	if err != nil {
		return 0, err
	}
	return length, nil
}
