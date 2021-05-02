package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandleOmikujiAPI(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	HandleOmikujiAPI(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Errorf("expected %d to eq %d", http.StatusOK, rw.StatusCode)
	}
}

func TestHandleOmikujiAPIHappyNewYear(t *testing.T) {
	getTime = func() time.Time { return time.Date(2018, time.January, 1, 0, 0, 0, 0, time.Local) }

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	HandleOmikujiAPI(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Errorf("expected %d to eq %d", http.StatusOK, rw.StatusCode)
	}

	b, _ := ioutil.ReadAll(rw.Body)
	actual := string(b)
	expected := "{\"result_code\":0,\"result\":\"daikichi\"}"

	if !strings.Contains(actual, expected) {
		t.Errorf("expected %s to eq %s", expected, actual)
	}
}
