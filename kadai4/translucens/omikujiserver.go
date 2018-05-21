package main

import (
	"log"
	"net/http"

	"github.com/translucens/dojo1/kadai4/translucens/omikuji"
)

func main() {
	http.HandleFunc("/", omikuji.HTTPHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}
}
