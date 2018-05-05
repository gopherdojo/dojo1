package main

import (
	"net/http"
	"fmt"
	"github.com/prometheus/common/log"
	"strconv"
	"io/ioutil"
	"sync"
	"os"
	"io"
)

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
	for i := 0; i < procs; i++ {
		wg.Add(1)

		from := subFileLen * i
		to := subFileLen * (i + 1)

		if i == procs-1 {
			to += remain
		}

		go func(from int, to int, i int) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", URL, nil)
			if err != nil {
				log.Fatal(err)
			}

			rangeHeader := "bytes=" + strconv.Itoa(from) + "-" + strconv.Itoa(to-1)
			req.Header.Add("Range", rangeHeader)
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Range: %v, %v bytes\n", rangeHeader, resp.ContentLength)
			defer resp.Body.Close()

			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			err = ioutil.WriteFile(strconv.Itoa(i), bytes, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}

			wg.Done()
		}(from, to, i)
	}
	wg.Wait()

	file, err := os.Create("gobook.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < procs; i++ {
		subFile, err := os.Open(fmt.Sprint(i))
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(file, subFile)

		subFile.Close()
	}
}
