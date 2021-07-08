package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/translucens/dojo1/kadai3/translucens/2-dler/network"
)

var (
	split    int
	savePath string
	urls     []string
)

func init() {
	const (
		defaultSplit    = 2
		usageSplit      = "specify download fragments count."
		defaultSavePath = "./download"
		usageSavePath   = "specify file destination."
	)

	flag.IntVar(&split, "n", defaultSplit, usageSplit)
	flag.StringVar(&savePath, "d", defaultSavePath, usageSavePath)
}

func main() {
	flag.Parse()
	urls = flag.Args()

	if savePath[len(savePath)-1] != '/' {
		savePath = savePath + "/"
	}

	savePathStat, err := os.Stat(savePath)
	if err != nil {

		if os.IsNotExist(err) {
			if err := os.Mkdir(savePath, 0755); err != nil {
				log.Fatalln(err.Error())
				return
			}
		} else {
			log.Fatalln(err.Error())
			return
		}
	} else if !savePathStat.IsDir() {
		fmt.Println("destination file exists")
	}

	for _, url := range urls {
		if err := network.Download(url, split, savePath); err != nil {
			log.Fatalln(err.Error())
		}
	}

}
