package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Result struct {
	Fortune string `json:"fortune"`
}

func omikuji() string {
	if isSanganichi(time.Now()) {
		return "大吉"
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(6)

	switch n + 1 {
	case 6:
		return "大吉"
	case 5:
	case 4:
		return "中吉"
	case 3:
	case 2:
		return "吉"
	}

	// case 1, write separately to avoid compile error.
	return "凶"
}

func isSanganichi(t time.Time) bool {
	if t.Month() == 1 && t.Day() <= 3 {
		return true
	}

	return false
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := &Result{Fortune: omikuji()}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(res); err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, buf.String())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
