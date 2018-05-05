package main

import (
	"net/http"
	"fmt"
	"strconv"
	"sync"
	"os"
	"io"
	"path"
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/prometheus/common/log"

)

const tempDir = "dlTmp"

var wg sync.WaitGroup

func main() {
	URL := "http://www.golang-book.com/public/pdf/gobook.pdf"
	procs := 5
	res, err := http.Head(URL)
	if err != nil {
		log.Fatal(err)
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		log.Fatal("this site doesn't support a range request")
	}
	len, err := strconv.Atoi(res.Header.Get("Content-Length"))
	fmt.Printf("total length: %d bytes\n", len)
	if err != nil {
		log.Fatal(err)
	}
	subFileLen := len / procs
	remain := len % procs

	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	eg, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < procs; i++ {
		i := i

		from := subFileLen * i
		to := subFileLen * (i + 1)

		if i == procs-1 {
			to += remain
		}

		eg.Go(func() error {
			return rangeRequest(ctx, from, to , i, URL)
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("gobook.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < procs; i++ {
		subFile, err := os.Open(path.Join(tempDir, fmt.Sprint(i)))
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(file, subFile)

		subFile.Close()
	}

	if err := os.RemoveAll(tempDir); err != nil {
		log.Fatal(err)
	}
}

func rangeRequest(ctx context.Context,from int, to int, i int, url string) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	rangeHeader := "bytes=" + strconv.Itoa(from) + "-" + strconv.Itoa(to-1)
	req.Header.Add("Range", rangeHeader)
	// errgroup.WithContext wraps context by calling context.WithCancel
	// cf. https://github.com/golang/sync/blob/master/errgroup/errgroup.go#L34
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("Range: %v, %v bytes\n", rangeHeader, resp.ContentLength)
	defer resp.Body.Close()

	file, err := os.OpenFile(path.Join(tempDir, fmt.Sprint(i)), os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, resp.Body)

	return nil
}
