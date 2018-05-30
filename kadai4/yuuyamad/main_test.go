package main

import (
	"testing"
	"time"
	"net/http/httptest"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"net/http"
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

	if rw.StatusCode != http.StatusOK { t.Fatal("unexpected status code") }
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil { t.Fatal("unexpected error") }
	const expected = "{\"omikuji\":\"大吉\"}"
	if s := string(b); !assert.JSONEq(t, s, expected) { t.Fatalf("unexpected response: %s", s) }


	//吉がでるseed
	date , _ := time.Parse(timeformat, "2018-02-03 14:10:00")
	setNow(date)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/omikuji", nil)
	GetOmikuji(w, r)
	rw = w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK { t.Fatal("unexpected status code") }
	b, err = ioutil.ReadAll(rw.Body)
	if err != nil { t.Fatal("unexpected error") }
	const expected2 = "{\"omikuji\":\"吉\"}"
	if s := string(b); !assert.JSONEq(t, s, expected2) { t.Fatalf("unexpected response: %s", s) }
}
