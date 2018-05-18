package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"strings"
)

func Test_Main(t *testing.T)  {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected tatus code")
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	const expected = "result"
	if !strings.Contains(string(b), expected) {
		t.Fatalf("unexpected response: %s", b)
	}

}