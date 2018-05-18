package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"strings"
	"time"
)

func Test_Main(t *testing.T)  {
	fakeDate()

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

	const expected = "大吉"
	if !strings.Contains(string(b), expected) {
		t.Fatalf("unexpected response: %s", b)
	}
}

func fakeDate()  {
	now = func() time.Time {
		return time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local )
	}
}