package main

import (
	"net/http"
	"fmt"
	"github.com/prometheus/common/log"
	"strconv"
	"io/ioutil"
	"sync"
)

var wg sync.WaitGroup

func main() {
	res, err := http.Head("http://www.golang-book.com/public/pdf/gobook.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Header.Get("Accept-Ranges"))
	if res.Header.Get("Accept-Ranges") != "bytes" {
		fmt.Println("this site doesn't support a range request")
	}
	l, err := strconv.Atoi(res.Header.Get("Content-Length"))
	fmt.Printf("total length: %d bytes\n", l)
	if err != nil {
		log.Fatal(err)
	}
	procs := 5
	clen := l / procs
	diff := l % procs
	body := make([]string, l+1)
	for i := 0; i < procs; i++ {
		wg.Add(1)

		from := clen * i
		to := clen * (i + 1)

		if i == procs-1 {
			to += diff // Add the remaining bytes in the last request
		}

		go func(min int, max int, i int) {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", "http://www.golang-book.com/public/pdf/gobook.pdf", nil)
			rangeHeader := ""
			if i != procs-1 {
				rangeHeader = "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max-1)
			} else {
				rangeHeader = "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max)
			}
			req.Header.Add("Range", rangeHeader)
			resp, _ := client.Do(req)
			fmt.Println(resp.ContentLength)
			defer resp.Body.Close()
			reader, _ := ioutil.ReadAll(resp.Body)
			body[i] = string(reader)
			ioutil.WriteFile(strconv.Itoa(i), []byte(string(body[i])), 0x777)
			wg.Done()
		}(from, to, i)
	}
	wg.Wait()
	//ioutil.WriteFile("gobook.pdf", []byte(body), 0x777)
}
