package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
	"flag"
	"fmt"
	"os"
)


const (
	// Exit codes are int values that represent an exit code for a particular error.
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
	// DefaultPort default for http server listen port
 	DefaultPort = 8080
)

var (
	logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
)

// OmikujiResponse omikuji-api response
type OmikujiResponse struct {
	ResultCode int    `json:"result_code"`
	Result     string `json:"result"`
}

var getTime = func() time.Time {
	return time.Now()
}

func omikuji() (int, string) {
	t := getTime()
	if t.Month() == time.January && t.Day() >= 1 && t.Day() <= 3 {
		return 0, "daikichi"
	}
	rand.Seed(t.UnixNano())
	s := rand.Intn(6) + 1

	switch s {
	case 1:
		return s, "kyo"
	case 2, 3:
		return s, "kichi"
	case 4, 5:
		return s, "chukichi"
	case 6:
		return s, "daikichi"
	default:
		return s, ""
	}
}

// HandleOmikujiAPI omikuji api
func HandleOmikujiAPI(w http.ResponseWriter, r *http.Request) {
	logger.Println("request: ", r.URL)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var v OmikujiResponse
	v.ResultCode, v.Result = omikuji()

	if err := json.NewEncoder(w).Encode(v); err != nil {
		logger.Println("error:", err)
	}
}

func main() {
	var (
		port    int
		version bool
	)

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.IntVar(&port, "port", DefaultPort, "port number")
	flags.IntVar(&port, "p", DefaultPort, "port number(Short)")
	flags.BoolVar(&version, "version", false, "print version information")
	if err := flags.Parse(os.Args[1:]); err != nil {
		os.Exit(ExitCodeError)
	}

	if version {
		fmt.Printf("%s version %s\n", Name, Version)
		os.Exit(ExitCodeOK)
	}

	http.HandleFunc("/", HandleOmikujiAPI)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
