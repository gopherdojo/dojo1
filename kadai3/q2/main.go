package main

import (
	"flag"
	"net/http"
	"log"
	"os"
	"strconv"
	"fmt"
	"io/ioutil"
	"sync"
	"strings"
	"context"
	"time"
	"./modules/httpHeader"
	"./modules/split"
)

var wg sync.WaitGroup

func main() {

	// url の取得
	var (
		url string
		limit int
	)
	flag.StringVar(&url, "url", "", "URL to download file")
	flag.StringVar(&url, "u", "", "URL to download file")
	flag.IntVar(&limit, "limit", 1, "Number of concurrent downloads.")
	flag.IntVar(&limit, "l", 1, "Number of concurrent downloads.")
	flag.Parse()

	// 入力値チェック
	if limit < 1 {
		fmt.Println("Limit number must be a positive number.")
		os.Exit(1)
	}

	// ヘッダの取得
	length, _ := httpHeader.GetLength(url)

	// ダウンロードする領域区切り毎の配列を生成
	limits := split.CreateArray(length, limit)

	// 各ダウンロード結果を保管する場所を用意
	body := make([]string, limit)

	// 分割ダウンロード
	var prev int = 0

	// キャンセル用のチャンネル
	//stopCh := make(chan struct{})

	for i, v := range limits {
		wg.Add(1)

		//go loop(stopCh, &wg)

		go func(min, max, i int, url string){
			client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatal(err)
			}

			rangeHeader := "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max - 1)
			req.Header.Add("Range", rangeHeader)
			resp,_ := client.Do(req)
			defer resp.Body.Close()
			reader, _ := ioutil.ReadAll(resp.Body)
			body[i] = string(reader)
			wg.Done()
		}(prev, v, i, url)
		wg.Wait()
		prev = v
	}
	// マージ
	mergedData := strings.Join(body, "")
	// 書き込み
	ioutil.WriteFile("result.jpg", []byte(mergedData), 0x777)

	return
}

//func loop(stopCh chan struct{}, wg *sync.WaitGroup) {
//	defer func() { wg.Done() }()
//
//	for {
//		println("(goroutine) loop...")
//		time.Sleep(1 * time.Second)
//
//		select {
//		case <- stopCh:
//			println("(goroutine) stop request received.")
//			return
//		default:
//			return
//		}
//	}
//}