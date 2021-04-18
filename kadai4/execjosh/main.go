package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gopherdojo/dojo1/kadai4/execjosh/internal/jinja"
)

func main() {
	j := jinja.New(time.Now().UnixNano())
	handler := func(w http.ResponseWriter, r *http.Request) {
		b := j.GetBlessing(time.Now())
		mkj := &jinja.Omikuji{Blessing: b}
		json.NewEncoder(w).Encode(mkj)
	}

	http.HandleFunc("/omikuji", handler)
	http.ListenAndServe(":8080", nil)
}
