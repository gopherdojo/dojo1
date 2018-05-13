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
)

const LIMIT=10

type Range struct {
	low    int
	high   int
	worker int
}

func main() {

	err := download()
	if err != nil {
		logError(err)
	}

	//分割ダウンロードしたファイルを結合する
	fh, err := os.Create("hoge.png")
	if err != nil {
		logError(err)
	}
	defer fh.Close()

	for j:=0; j < LIMIT; j++ {
		f := fmt.Sprintf("%s.%d", "hoge.png", j)
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
func download() error{
	url := os.Args[1]

	// contents Headerを取得する
	res, err := http.Head(url)
	if err != nil {
		logError(err)
	}

	maps := res.Header
	length, err := strconv.Atoi(maps["Content-Length"][0])

	len_sub := length/LIMIT
	diff := length%LIMIT

	// errorGroup
	grp, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()


	for i:=0; i < LIMIT; i++ {

		min := len_sub * i
		max := len_sub * (i + 1)

		r := Range{}
		r.low = min
		r.high = max
		r.worker = i

		if(i == (LIMIT - 1)){
			max += diff
		}
		// execute get request
		grp.Go(func() error {
			return requests(ctx, url, r)
		})

		fmt.Println(i)
	}
	if err := grp.Wait(); err != nil {
		return err
	}
	return nil
}

func requests(ctx context.Context, url string, r Range) error {
	body := make([]string, 11)
	client := &http.Client{}
	req , err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	range_header := "bytes=" + strconv.Itoa(r.low) + "-" + strconv.Itoa(r.high-1)
	req.Header.Add("Range", range_header)
	resp,err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	reader, err := ioutil.ReadAll(resp.Body)
	body[r.worker] = string(reader)
	fmt.Println(r.worker)
	ioutil.WriteFile("hoge.png." + strconv.Itoa(r.worker), []byte(string(body[r.worker])), 0x777)

	return nil
}
