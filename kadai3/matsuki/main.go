package main

import (
	"flag"
	"log"
	"os"

	"./pdl"
	"github.com/hashicorp/logutils"
)

var (
	trgURL    string // URL of download file
	isSingle  bool
	isVerbose bool
	procsNum  int64
)

func init() {
	flag.StringVar(&trgURL, "url", "", "url of download file")
	flag.Int64Var(&procsNum, "n", 0, "procs number")
	flag.BoolVar(&isSingle, "s", false, "single download or not (default: false =parallel download)")
	flag.BoolVar(&isVerbose, "v", false, "verbose")
	flag.Parse()
}

func main() {
	logLevelStr := ""
	if isVerbose {
		logLevelStr = "DEBUG"
	} else {
		logLevelStr = "WARN"
	}

	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(logLevelStr),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	var procs uint
	if isSingle == false && procsNum > 1 {
		procs = uint(procsNum)
	} else if isSingle == false && procsNum == 0 {
		procs = uint(0)
	} else {
		procs = 1
	}

	p, err := pdl.NewClient(trgURL, isSingle == false, procs)
	if err != nil {
		flag.PrintDefaults()
		log.Fatal(err)
	}

	p.Download()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("downloaded file of %s", p.Filename())
}
