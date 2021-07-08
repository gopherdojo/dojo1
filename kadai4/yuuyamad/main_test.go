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
	heijitu , _ := time.Parse(timeformat, "2018-02-03 14:10:00")
	heijitu2 , _ := time.Parse(timeformat, "2018-03-03 18:11:11")

	cases := []struct {
		date      time.Time
		responce string
	}{
		{date: syogatsu, responce: "{\"omikuji\":\"大吉\"}"},
		{date: heijitu, responce: "{\"omikuji\":\"吉\"}"},
		{date: heijitu2, responce: "{\"omikuji\":\"中吉\"}"},

	}

	for _, c := range cases {
		setNow(c.date)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/api/v1/omikuji", nil)
		GetOmikuji(w, r)
		rw := w.Result()
		defer rw.Body.Close()

		if rw.StatusCode != http.StatusOK {
			t.Fatal("unexpected status code")
		}
		b, err := ioutil.ReadAll(rw.Body)
		if err != nil {
			t.Fatal("unexpected error")
		}

		if s := string(b); !assert.JSONEq(t, s, c.responce) {
			t.Fatalf("unexpected response: %s", s)
		}

	}
}
