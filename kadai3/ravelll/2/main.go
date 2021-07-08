package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const (
	BatchByteSize = 500
)

func main() {
	contentURL := os.Args[1]
	os.Exit(download(contentURL))
}

func checkContentLength(contentURL string) (int64, error) {
	req, err := http.Head(contentURL)
	if err != nil {
		return 0, err
	}

	return req.ContentLength, nil
}

func rangeHeaderVal(i int, partNum int) string {
	if i == partNum-1 {
		return fmt.Sprintf(
			"bytes=%d-", i*BatchByteSize)
	}

	return fmt.Sprintf(
		"bytes=%d-%d", i*BatchByteSize, (i+1)*BatchByteSize-1)
}

func request(contentURL string) (string, error) {
	partedBody := make([]string, 0, 20)

	contentLength, err := checkContentLength(contentURL)
	if err != nil {
		return "", err
	}

	partNum := int((contentLength / BatchByteSize) + 1)

	for i := 0; partNum > i; i++ {
		client := http.Client{}
		req, err := http.NewRequest("GET", contentURL, nil)
		req.Header.Add("Range", rangeHeaderVal(i, partNum))
		downloaded, err := client.Do(req)
		respBody, err := ioutil.ReadAll(downloaded.Body)
		if err != nil {
			return "", err
		}

		partedBody = append(partedBody, string(respBody))
	}

	var body string

	for i := 0; i < len(partedBody); i++ {
		body += partedBody[i]
	}

	return body, nil
}

func getFilename(contentURL string) (string, error) {
	parsed, err := url.Parse(contentURL)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	_, filename := filepath.Split(parsed.EscapedPath())
	return filename, nil
}

func download(contentURL string) int {
	if contentURL == "" {
		fmt.Println("Error: no url specified.")
		return 1
	}

	filename, filenameErr := getFilename(contentURL)
	if filenameErr != nil {
		return 1
	}

	responseBody, reqErr := request(contentURL)
	if reqErr != nil {
		return 1
	}

	output, createErr := os.Create(filename)
	defer output.Close()
	if createErr != nil {
		return 1
	}

	output.WriteString(responseBody)

	return 0
}
