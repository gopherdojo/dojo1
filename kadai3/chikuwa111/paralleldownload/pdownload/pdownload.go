package pdownload

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/sync/errgroup"
)

func checkAcceptRanges(res *http.Response) bool {
	acceptRanges := res.Header.Get("Accept-Ranges")
	return acceptRanges == "bytes"
}

func generateFileName(url string) string {
	slice := strings.Split(url, "/")
	return slice[len(slice)-1]
}

func generateRange(contentLength int64, parallelCount int) []string {
	ranges := []string{}
	onepart := contentLength / int64(parallelCount)
	var borderByte int64 // = 0
	for i := 0; i < parallelCount-1; i++ {
		ranges = append(ranges, fmt.Sprintf("bytes=%v-%v", borderByte, borderByte+onepart))
		borderByte += onepart + 1
	}
	ranges = append(ranges, fmt.Sprintf("bytes=%v-%v", borderByte, contentLength-1))
	return ranges
}

func download(url, tmpFileName, rangeStr string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", rangeStr)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	file, err := os.Create(tmpFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, res.Body)

	return nil
}

func parallelDownload(url, fileName string, ranges []string) error {
	var errgrp errgroup.Group
	for i, rangeStr := range ranges {
		tmpFileName := fmt.Sprintf("%v.%v", fileName, i)
		errgrp.Go(func() error {
			return download(url, tmpFileName, rangeStr)
		})
	}
	err := errgrp.Wait()
	return err
}

func bundleOneFile(fileName string, parallelCount int) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; i < parallelCount; i++ {
		tmpFileName := fmt.Sprintf("%v.%v", fileName, i)
		tmpFile, err := os.Open(tmpFileName)
		if err != nil {
			return err
		}

		io.Copy(file, tmpFile)

		tmpFile.Close()

		err = os.Remove(tmpFileName)
		if err != nil {
			return err
		}
	}
	return nil
}

// Run executes parallel download.
func Run(url string, parallelCount int) error {
	res, err := http.Head(url)
	if err != nil {
		return err
	}

	if !checkAcceptRanges(res) {
		return errors.New("Cannot download parallel")
	}

	contentLength := res.ContentLength
	if contentLength <= 0 {
		return errors.New("Invalid Content-Length: " + string(contentLength))
	}

	fileName := generateFileName(url)
	ranges := generateRange(contentLength, parallelCount)

	if err := parallelDownload(url, fileName, ranges); err != nil {
		return err
	}

	if err := bundleOneFile(fileName, parallelCount); err != nil {
		return err
	}

	return nil
}
