package gurl

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/sync/errgroup"
)

// Client gurl client
type Client struct {
	Parallel int
	Output   string
}

// Content actual content
type Content struct {
	Name   string
	Length int
}

// NewClient constractor for Client
func NewClient(parallel int, output string) *Client {
	return &Client{
		Parallel: parallel,
		Output:   output,
	}
}

func rangeGet(ctx context.Context, url string, s, e, i int, tempFiles []*os.File) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", s, e))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	reader, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	tempFile, err := ioutil.TempFile("./", "temp")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(tempFile.Name(), reader, 0644); err != nil {
		return err
	}
	tempFiles[i] = tempFile

	return nil
}

// Get content of url
func (c *Client) Get(url string) error {
	resp, err := http.Head(url)
	if err != nil {
		return err
	}

	contentLength, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	chunkSize := contentLength / c.Parallel
	surplus := contentLength % c.Parallel

	eg, ctx := errgroup.WithContext(context.TODO())
	tempFiles := make([]*os.File, c.Parallel)

	for p := 0; p < c.Parallel; p++ {
		s := p * chunkSize
		e := s + (chunkSize - 1)
		if p == c.Parallel-1 {
			e += surplus
		}

		i := p
		eg.Go(func() error {
			return rangeGet(ctx, url, s, e, i, tempFiles)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	tempFilesReaders := make([]io.Reader, c.Parallel)
	for i, f := range tempFiles {
		tempFilesReaders[i], err = os.Open(f.Name())
		if err != nil {
			return err
		}
	}

	reader := io.MultiReader(tempFilesReaders...)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(c.Output, b, 0644); err != nil {
		return err
	}

	for _, f := range tempFiles {
		os.Remove(f.Name())
	}

	return nil
}
