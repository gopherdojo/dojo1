package omikuji_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/xlune/dojo1/kadai4/xlune/omikuji"
)

var handlerTests = []struct {
	date    string
	fortune omikuji.Fortune
}{
	{"2017-01-01 00:00:00", omikuji.Daikichi},
	{"2017-01-01 23:59:59", omikuji.Daikichi},
	{"2017-01-02 00:00:00", omikuji.Daikichi},
	{"2017-01-02 23:59:59", omikuji.Daikichi},
	{"2017-01-03 00:00:00", omikuji.Daikichi},
	{"2017-01-03 23:59:59", omikuji.Daikichi},
	{"2018-04-01 01:23:45", omikuji.Chukichi},
	{"2019-08-08 22:23:45", omikuji.Kyo},
	{"2020-10-06 10:10:10", omikuji.Daikichi},
	{"2021-12-22 11:11:09", omikuji.Kichi},
}

func TestHandler(t *testing.T) {
	for _, h := range handlerTests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		q := r.URL.Query()
		q.Add("date", h.date)
		r.URL.RawQuery = q.Encode()
		omikuji.Handler(w, r)
		rw := w.Result()
		defer rw.Body.Close()
		if rw.StatusCode != http.StatusOK {
			t.Fatal("unexpected status code")
		}
		var res omikuji.Result
		dec := json.NewDecoder(rw.Body)
		if err := dec.Decode(&res); err != nil {
			t.Fatal("json decode error")
		}
		if res.Type != h.fortune {
			t.Fatal("unexpected error")
		}
	}
}
