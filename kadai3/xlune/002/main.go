package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xlune/dojo1/kadai3/xlune/002/client"
	"golang.org/x/sync/errgroup"
)

var (
	splitCount int
)

func init() {
	flag.IntVar(&splitCount, "d", 3, "分割数")
}

func main() {
	flag.Parse()

	sourcePath := flag.Arg(0)

	url, err := client.GetSourceURL(sourcePath)
	if err != nil {
		fmt.Printf("URLチェックに失敗しました...(%s)\n", err.Error())
		os.Exit(1)
	}

	info, err := client.GetInfo(url)
	if err != nil {
		fmt.Printf("ファイル情報取得に失敗しました...(%s)\n", err.Error())
		os.Exit(1)
	}

	err = download(info, splitCount)
	if err != nil {
		fmt.Printf("ダウンロードに失敗しました...(%s)\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("ダウンロード完了しました！ (%s)\n", info.Name)
}

func download(info client.TargetInfo, count int) error {

	// レンジ取得未対応かファイルサイズが小さいものは分割しない
	if !info.IsRanges || info.Size < int64(count)*client.UnitSize {
		count = 1
	}

	// 分割サイズリスト取得
	ranges, err := client.MakeDownloadRange(info.Size, count)
	if err != nil {
		return err
	}

	// アイテム情報を生成して分割ダウンロード
	eg := errgroup.Group{}
	items := []*client.DownloadItem{}
	for index, r := range ranges {
		item := &client.DownloadItem{
			URL:        info.URL,
			RangeIndex: index,
			Range:      r,
		}
		eg.Go(func() error {
			return client.DownloadByItem(item)
		})
		items = append(items, item)
	}

	// 処理を待ってエラー判定
	err = eg.Wait()
	if err != nil {
		return err
	}

	// 書き込み先を新規でオープン
	file, err := os.OpenFile(info.Name, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// 書き込み
	for _, item := range items {
		_, err := file.WriteAt(item.Body, item.Range.Start)
		if err != nil {
			return err
		}
	}

	return nil
}
