package main

import (
	"flag"
	"log"
	"path/filepath"

	"./downloader"
)

func main() {
	parallel := flag.Int("p", 6, "並列度")
	dir := flag.String("d", "./", "ディレクトリ")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("ダウンロードURLを指定してください")
	}

	url := args[0]
	_, name := filepath.Split(url)
	path, err := filepath.Abs(filepath.Join(*dir, name))
	if err != nil {
		log.Fatal(err)
	}

	err = downloader.Run(url, path, *parallel)
	if err != nil {
		log.Fatal(err)
	}
}
