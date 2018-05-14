package main

import (
	"os"
	"net/http"
	"strconv"
	"io/ioutil"
	"fmt"
	"io"
	"golang.org/x/sync/errgroup"
	"context"
	"time"
	"flag"
)

type Range struct {
	low    int
	high   int
	worker int
}

type Downloader struct {
	procs int
	filename string
	url string
}

func main() {

	var (
		procs = flag.Int("p", 10, "split ratio to download file")
		output = flag.String("o", "", "output filename")
	)

	flag.Parse()
	args := flag.Args()

	downloder := Downloader{}
	downloder.procs = *procs
	downloder.filename = *output
	downloder.url = args[0]


	err := downloder.download()
	if err != nil {
		logError(err)
	}

	//分割ダウンロードしたファイルを結合する
	fh, err := os.Create(downloder.filename)
	if err != nil {
		logError(err)
	}
	defer fh.Close()

	for j:=0; j < downloder.procs; j++ {
		f := fmt.Sprintf("%s.%d", downloder.filename, j)
		subfp, err := os.Open(f)
		if err != nil{
			logError(err)
		}

		io.Copy(fh, subfp)

		subfp.Close()
		if err := os.Remove(f); err != nil {
			logError(err)
		}
	}
}
func (d *Downloader)download() error{

	// contents Headerを取得する
	res, err := http.Head(d.url)
	if err != nil {
		return err
	}

	maps := res.Header
	length, err := strconv.Atoi(maps["Content-Length"][0])

	len_sub := length/d.procs
	diff := length%d.procs

	// errorGroup
	grp, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()


	for i:=0; i < d.procs; i++ {

		min := len_sub * i
		max := len_sub * (i + 1)

		r := Range{}
		r.low = min
		r.high = max
		r.worker = i

		if(i == (d.procs - 1)){
			max += diff
		}
		// execute get request
		grp.Go(func() error {
			return d.requests(ctx, r)
		})
	}
	if err := grp.Wait(); err != nil {
		return err
	}
	return nil
}

func (d *Downloader)requests(ctx context.Context, r Range) error {
	body := make([]string, 99)
	client := &http.Client{}
	req , err := http.NewRequest("GET", d.url, nil)
	if err != nil {
		return err
	}

	range_header := "bytes=" + strconv.Itoa(r.low) + "-" + strconv.Itoa(r.high-1)
	req.Header.Add("Range", range_header)

	errCh := make(chan error, 1)
	tmpfile := d.filename + "." + strconv.Itoa(r.worker)

	go func() {
		resp, err := client.Do(req)
		defer resp.Body.Close()

		reader, err := ioutil.ReadAll(resp.Body)
		body[r.worker] = string(reader)

		err = ioutil.WriteFile(tmpfile, []byte(string(body[r.worker])), 0x777)
		errCh <- err
	}()

	select {
	case err := <-errCh:
		fmt.Printf("requests: %s\n", err)
		if err != nil {
			return err
		}
	case <-ctx.Done():
		fmt.Printf("requests: %s\n", ctx.Err())
		os.Remove(tmpfile)
		<-errCh
		return ctx.Err()

	}
	return nil
}
