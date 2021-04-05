package omikuji

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

func contains(expecteds []string, actual string) bool {
	for _, expected := range expecteds {
		if expected == actual {
			return true
		}
	}
	return false
}

func Testランダムに結果を返す(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(Handler))
	defer s.Close()

	for i := 0; i < 100; i++ {
		res, err := http.Get(s.URL)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("unexpected status code: %d", res.StatusCode)
		}

		contentType := res.Header.Get("Content-Type")
		if contentType != "application/json; charset=UTF-8" {
			t.Errorf("unexpected content type: %s", contentType)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		var result Result
		if err := json.Unmarshal(body, &result); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !(contains(結果, result.Result)) {
			t.Errorf("unexpected contents: %s", result.Result)
		}

		if _, err := time.Parse(time.RFC3339, result.Date); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

// refer: https://blog.questionable.services/article/testing-http-handlers-go/
func Test三が日は大吉を返す(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	rand.Seed(time.Now().UnixNano())

	days := []int{1, 2, 3}
	for _, day := range days {
		for i := 0; i < 100; i++ {
			date := time.Date(2018, time.January, day, rand.Intn(24), rand.Intn(60), rand.Intn(60), rand.Intn(1e9), jst)

			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			ctx := context.WithValue(req.Context(), datetimeContextKey, date)
			handler.ServeHTTP(recorder, req.WithContext(ctx))

			if recorder.Code != http.StatusOK {
				t.Errorf("unexpected status code: %d", recorder.Code)
			}
			contentType := recorder.Header().Get("Content-Type")
			if contentType != "application/json; charset=UTF-8" {
				t.Errorf("unexpected content type: %s", contentType)
			}

			body, err := ioutil.ReadAll(recorder.Body)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			var result Result
			if err := json.Unmarshal(body, &result); err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !(contains(結果, result.Result)) {
				t.Errorf("unexpected contents: %s", result.Result)
			}

			if result.Date != date.Format(time.RFC3339) {
				t.Errorf("unexpected date: %v", result.Date)
			}
		}
	}
}
