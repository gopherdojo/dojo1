package main

import (
	"testing"
	"time"
	"net/http/httptest"
	"io/ioutil"
	"fmt"
)

const timeformat = "2006-01-02 15:04:06"

func setNow(t time.Time) {
	timeNowFunc = func() time.Time { return t }
}


func TestGetOmikujiain(t *testing.T) {
	syogatsu , _ := time.Parse(timeformat, "2018-01-03 14:10:00")
	setNow(syogatsu)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/omikuji", nil)
	GetOmikuji(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	//if rw.StatusCode != http.StatusOK { t.Fatal("unexpected status code") }
	b, err := ioutil.ReadAll(rw.Body)
	fmt.Println(string(b))
	if err != nil { t.Fatal("unexpected error") }
	const expected = "{omikuji:\"大吉\"}"
	if s := string(b); s != expected { t.Fatalf("unexpected response: %s", s) }



}
