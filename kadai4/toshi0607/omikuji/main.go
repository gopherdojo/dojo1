package main

import (
	"net/http"
	"time"
	"math/rand"
	"encoding/json"
	"log"
)

var now = time.Now

func init() {
	t := now().UnixNano()
	rand.Seed(t)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

type omikuji struct {
	Result string `json:"result"`
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	result := drawOmikuji()
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Println("Error: ", err)
	}
}

func drawOmikuji() omikuji {
	if isShogatsu() {
		return omikuji{Result: "大吉"}
	}

	var result string
	switch rand.Intn(7) {
	case 6:
		result = "大吉"
	case 5, 4:
		result = "吉"
	case 3, 2, 1:
		result = "中吉"
	case 0:
		result = "凶"
	}
	return omikuji{Result: result}
}

func isShogatsu() bool {
	now := now()
	return now.Month() == 1 && now.Day() <= 3
}
