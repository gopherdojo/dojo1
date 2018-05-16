package network

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	filenameRegexp = regexp.MustCompile(`([^/]+?)$`)
	tempDir        string
)

func init() {
	var err error
	tempDir, err = ioutil.TempDir("", "dler")
	if err != nil {
		panic(err)
	}
}

type downloadResult struct {
	downloadedBytes int64
	err             error
}

// GetFileSize returns file size of indicated by URL
func GetFileSize(url string) (int64, error) {

	res, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.New("this server does not support partial requests")
	}
	if res.StatusCode != 200 {
		return 0, errors.New("File " + url + " is not available. HTTP: " + strconv.Itoa(res.StatusCode))
	}

	return res.ContentLength, nil
}

func cutFileName(path string) string {
	if path[len(path)-1] == '/' {
		path = path[0 : len(path)-1]
	}

	return filenameRegexp.FindString(path)
}

// Download downloads specified files from URL
func Download(rawurl string, fragments int, savepath string) error {
	defer CleanTempDir()
	startAt := time.Now()

	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	filename := cutFileName(parsedURL.Path)
	fileSize, err := GetFileSize(rawurl)
	if err != nil {
		return err
	}

	fragmentIndices := make([]int64, fragments+1)
	for i := range fragmentIndices {
		fragmentIndices[i] = fileSize * int64(i) / int64(fragments)
	}

	fragmentPaths := make([]string, fragments)
	for i := range fragmentPaths {
		fragmentPaths[i] = fmt.Sprintf("%s/%s.%d.tmp", tempDir, filename, i)
	}

	errg := errgroup.Group{}

	for i := range fragmentPaths {
		from := fragmentIndices[i]
		to := fragmentIndices[i+1] - 1
		fragmentPath := fragmentPaths[i]

		errg.Go(
			func() error {
				return DownloadFragment(rawurl, fragmentPath, from, to)
			})
	}
	if err := errg.Wait(); err != nil {
		return err
	}

	spentTime := float32((time.Now().Sub(startAt))/time.Millisecond) / 1000
	Mbps := float32(fileSize) * 8 / spentTime / 1000000
	fmt.Printf("Total: %d bytes (%.2f sec. / %.3f Mbps)\n", fileSize, spentTime, Mbps)
	return Concatenate(fragmentPaths, savepath+filename)
}

// DownloadFragment download specified URL in the range
// returns downloaded bytes
func DownloadFragment(url, filepath string, from, to int64) error {

	fmt.Printf("Downloading %d - %d: %s\n", from, to, filepath)
	startAt := time.Now()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if from < to {
		req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", from, to))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	bytes, err := WriteFile(filepath, res.Body, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)

	if bytes != to-from+1 {
		return fmt.Errorf("file size doen not match with expected %d byte; actual %d byte", to-from+1, bytes)
	}

	spentTime := float32((time.Now().Sub(startAt))/time.Millisecond) / 1000
	Mbps := float32(bytes) * 8 / spentTime / 1000000
	fmt.Printf("Downloaded %d - %d (%.2f sec. / %.3f Mbps)\n", from, to, spentTime, Mbps)
	return nil
}

// WriteFile writes content to filepath
func WriteFile(filepath string, content io.Reader, fileflag int) (int64, error) {

	ofd, err := os.OpenFile(filepath, fileflag, 0755)
	if err != nil {
		return 0, err
	}
	defer ofd.Close()

	buffered := bufio.NewWriter(ofd)
	defer buffered.Flush()

	return io.Copy(buffered, content)
}

// Concatenate combines source files into one destination file
func Concatenate(srcs []string, dst string) error {

	for _, src := range srcs {
		reader, err := os.Open(src)
		if err != nil {
			return err
		}
		_, err = WriteFile(dst, reader, os.O_WRONLY|os.O_APPEND|os.O_CREATE)
		if err != nil {
			return err
		}
	}

	return nil
}

// CleanTempDir removes all temp files
func CleanTempDir() {
	os.RemoveAll(tempDir)
}
