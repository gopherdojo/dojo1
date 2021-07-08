package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"sort"

	"golang.org/x/sync/errgroup"
)

const chunkFileDir = "chunks"

func DoParallel(urlPath, outputFilePath string, procs int) error {
	chunkFilePaths, err := downloadAsChunkFiles(urlPath, procs)
	if err != nil {
		return err
	}

	if err := concatChunkFiles(chunkFilePaths, outputFilePath); err != nil {
		return err
	}

	return nil
}

func downloadAsChunkFiles(urlPath string, procs int) ([]string, error) {
	contentLength, err := fetchContentLength(urlPath)
	if err != nil {
		return nil, err
	}

	rangeHeaders := generateRangeHeaders(contentLength, procs)
	chunkFilePathChan := make(chan string, 100)
	var eg errgroup.Group
	for i, rangeHeader := range rangeHeaders {
		eg.Go(func(u, r string, i int) func() error {
			return func() error {
				return downloadChunk(u, r, i, chunkFilePathChan)
			}
		}(urlPath, rangeHeader, i))
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	close(chunkFilePathChan)

	var chunkFilePaths []string
	for chunkFilePath := range chunkFilePathChan {
		chunkFilePaths = append(chunkFilePaths, chunkFilePath)
	}

	sort.Strings(chunkFilePaths)
	return chunkFilePaths, nil
}

func concatChunkFiles(chunkFilePaths []string, outputFilePath string) error {
	dir := path.Dir(outputFilePath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	resultFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}

	for _, chunkFilePath := range chunkFilePaths {
		subfp, err := os.Open(chunkFilePath)
		if err != nil {
			return err
		}
		io.Copy(resultFile, subfp)
		if err := os.Remove(chunkFilePath); err != nil {
			return err
		}
	}

	if err := os.Remove(chunkFileDir); err != nil {
		return err
	}

	return nil
}

func downloadChunk(urlPath, rangeHeader string, index int, filePathChan chan<- string) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Range", rangeHeader)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	chunkFilePath := createChunkFilePath(chunkFileDir, index)
	if err := os.MkdirAll(chunkFileDir, 0755); err != nil {
		return err
	}

	file, err := os.Create(chunkFilePath)
	defer file.Close()
	if err != nil {
		return err
	}

	io.Copy(file, resp.Body)
	filePathChan <- chunkFilePath
	return nil
}

func generateRangeHeaders(contentLength, splitNum int) (rangeHeaders []string) {
	bytesPerRange := contentLength / splitNum
	startByte := 0
	for i := 0; i < (splitNum - 1); i++ {
		rangeHeaders = append(rangeHeaders, fmt.Sprintf("bytes=%d-%d", startByte, startByte+bytesPerRange-1))
		startByte += bytesPerRange
	}

	rangeHeaders = append(rangeHeaders, fmt.Sprintf("bytes=%d-%d", startByte, contentLength-1))
	return
}

func createChunkFilePath(dir string, index int) string {
	return path.Join(dir, fmt.Sprintf("%04d_download", index))
}

func fetchContentLength(urlPath string) (int, error) {
	client := &http.Client{}
	req, err := http.NewRequest("HEAD", urlPath, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	contentLengthHeader := resp.Header.Get("Content-Length")
	contentLength, err := strconv.Atoi(contentLengthHeader)
	if err != nil {
		return 0, err
	}
	return contentLength, nil
}
