package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"math/rand"
	"io/ioutil"
	"time"
	"fmt"
)

func Test_Normal(t *testing.T) {
	t.Parallel()
	tests := []string{"大吉", "吉", "中吉", "凶"}
	for _, te := range tests {
		te = te
		count := 0
		for {
			if count == 30 {
				t.Fatalf("%s doesn't appear in %d requests", te, count)
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)

			s := Server{config: DefaultConfig()}
			rand.Seed(s.config.Now().UnixNano())
			s.Handler().ServeHTTP(w, r)

			rw := w.Result()
			defer rw.Body.Close()

			if rw.StatusCode != http.StatusOK {
				t.Fatalf("unexpected tatus code: %d", rw.StatusCode)
			}

			b, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			re := fmt.Sprintf("{\"result\":\"%s\"}\n", te)
			if string(b) == re {
				break
			}
			count++
		}

	}



}

func Test_Shogatsu(t *testing.T) {
	t.Parallel()
	tests := []struct{
		t time.Time
		result string
	}{
		{
			t: time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local ),
			result: "大吉",
		},
		{
			t: time.Date(2018, 1, 2, 3, 4, 0, 0, time.Local ),
			result: "大吉",
		},
		{
			t: time.Date(2018, 1, 3, 23, 59, 59, 0, time.Local ),
			result: "大吉",
		},
	}

	for _, te := range tests{
		s := Server{config: &Config{
			Now: func() time.Time { return te.t },
			Port: "8080",
		}}
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		rand.Seed(s.config.Now().UnixNano())
		s.Handler().ServeHTTP(w, r)

		rw := w.Result()
		defer rw.Body.Close()

		if rw.StatusCode != http.StatusOK {
			t.Fatalf("unexpected tatus code: %d", rw.StatusCode)
		}

		b, err := ioutil.ReadAll(rw.Body)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		re := fmt.Sprintf("{\"result\":\"%s\"}\n", te.result)
		if string(b) != re {
			t.Fatalf("want: %s, got: %s", re, string(b))
		}
	}
}
