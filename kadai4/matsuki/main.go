package main

import (
	"time"

	"github.com/gopherdojo/dojo1/kadai4/matsuki/lib"
)

var verbos bool

func main() {
	lib.Verbos = true

	f := lib.NewFortune(time.Now())
	// server
	s := lib.NewServer(f)
	s.ListenAndServe()
}
