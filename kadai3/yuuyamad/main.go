package main

import (
	"os"
	"net/http"
	"strconv"
	"sync"
	"io/ioutil"
	"fmt"
	"io"
)

func main() {

	url := os.Args[1]

	// contents Headerを取得する
	res, err := http.Head(url)
	if err != nil {
		logError(err)
	}

	maps := res.Header
	length, err := strconv.Atoi(maps["Content-Length"][0])

	limit := 10

	len_sub := length/limit
	diff := length%limit
	body := make([]string, 11)

	wg := &sync.WaitGroup{}

	for i:=0; i < limit; i++ {
		wg.Add(1)

		min := len_sub * i
		max := len_sub * (i + 1)

		if(i == limit - 1){
			max += diff
		}

		go func(min int, max int, i int) {
			client := &http.Client{}
			req , err := http.NewRequest("GET", url, nil)
			if err != nil {
				logError(err)
			}
			range_header := "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max-1)
			req.Header.Add("Range", range_header)
			resp,err := client.Do(req)
			if err != nil {
				logError(err)
			}
			defer resp.Body.Close()


			reader, err := ioutil.ReadAll(resp.Body)
			body[i] = string(reader)
			ioutil.WriteFile("hoge.png." + strconv.Itoa(i), []byte(string(body[i])), 0x777)

			wg.Done()


		}(min, max, i)
	}
	wg.Wait()

	//分割ダウンロードしたファイルを結合する
	fh, err := os.Create("hoge.png")
	if err != nil {
		logError(err)
	}
	defer fh.Close()

	for j:=0; j < limit; j++ {
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
