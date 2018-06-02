package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"./omikuji"
)

func main() {
	port := 8080
	if len(os.Args) > 1 {
		p, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		port = p
	}

	mux := http.NewServeMux()
	mux.Handle("/", omikuji.AddCurrentDateTime(http.HandlerFunc(omikuji.Handler)))
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
