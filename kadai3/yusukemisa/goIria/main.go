package main

import (
	"log"
	"os"

	"github.com/yusukemisa/goIria/iria"
)

//RFC 7233 — HTTP/1.1: Range Requests
//go run main.go  https://beauty.hotpepper.jp/CSP/c_common/ALL/IMG/cam_cm_327_98.jpg
func main() {
	downloader, err := iria.New(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := downloader.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
