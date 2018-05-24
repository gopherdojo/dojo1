package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"time"
)

func TestFortune(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	fortune(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK { t.Fatal("unexpected status code") }
	_, err := ioutil.ReadAll(rw.Body)
	if err != nil { t.Fatal("unexpected error") }
}

func TestGetLuckNormal(t *testing.T) {
	luck := getLuck(time.Now())
	if luck == "Unknown Number" {
		t.Error("An unexpected random number was generated")
	}
}

func TestGetLuckNewYear(t *testing.T) {
	local, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	luck := getLuck(time.Date(2018, time.January, 1, 0, 0, 0, 0, local ))
	if luck != "大吉" {
		t.Error("Daikichi has to come out when I pulled it in the New Year")
	}
}