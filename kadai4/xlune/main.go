package main

import (
	"net/http"

	"github.com/xlune/dojo1/kadai4/xlune/omikuji"
)

func main() {
	http.HandleFunc("/", omikuji.Handler)
	http.ListenAndServe(":8080", nil)
}
