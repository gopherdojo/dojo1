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
	t := time.Now().UnixNano()
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

	var result string
	now := now()
	if now.Month() == 1 && ( 1 <= now.Day() && now.Day() <= 3) {
		result = "大吉"
	} else {
		s := rand.Intn(7)

		switch s {
		case 6:
			result = "大吉"
		case 5, 4:
			result = "吉"
		case 3, 2, 1:
			result = "中吉"
		case 0:
			result = "凶"
		}
	}

	r := omikuji{Result: result}
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println("Error: ", err)
	}
}
