package omikuji

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_omikujiNewYear(t *testing.T) {

	nowFunc = func() time.Time {
		return time.Date(2018, time.January, 1, 0, 0, 0, 0, time.Local)
	}

	result := omikuji()
	if result != "大吉！" {
		t.Errorf("omikuji() = %s", result)
	}

}

func Test_omikujiNewYear2(t *testing.T) {

	nowFunc = func() time.Time {
		return time.Date(2018, time.January, 3, 23, 59, 59, 0, time.Local)
	}

	result := omikuji()
	if result != "大吉！" {
		t.Errorf("omikuji() = %s", result)
	}

}

func Test_omikujiNotNewYear(t *testing.T) {

	nowFunc = func() time.Time {
		return time.Date(2018, time.January, 4, 0, 0, 0, 0, time.Local)
	}

	switch result := omikuji(); result {
	case "大吉", "凶", "吉", "末吉", "末小吉", "半吉", "小吉":
		t.Logf("omikuji() = %s", result)
	default:
		t.Errorf("omikuji() = %s", result)
	}
}

func Test_omikujiNotNewYear2(t *testing.T) {

	nowFunc = func() time.Time {
		return time.Date(2017, time.December, 31, 23, 59, 59, 0, time.Local)
	}

	switch result := omikuji(); result {
	case "大吉", "凶", "吉", "末吉", "末小吉", "半吉", "小吉":
		t.Logf("omikuji() = %s", result)
	default:
		t.Errorf("omikuji() = %s", result)
	}
}

func TestHTTPHandler(t *testing.T) {

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	HTTPHandler(w, req)
	result := w.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code: %d", result.StatusCode)
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("Cannot read response body ", err)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(body))
	var res Response
	err = decoder.Decode(&res)
	if err != nil {
		t.Fatal("Decode error: ", err)
	}

	switch res.Result {
	case "大吉", "凶", "吉", "末吉", "末小吉", "半吉", "小吉":
		t.Logf("result = %s", res)
	default:
		t.Errorf("result = %s", res)
	}
}

func testResultCount(t *testing.T, actualCount, expectedCount int) {
	t.Helper()

	diff := 100

	if actualCount < expectedCount-diff {
		t.Errorf("The actual number is too small, actual: %d, expected: %d", actualCount, expectedCount)
	} else if expectedCount+diff < actualCount {
		t.Errorf("The actual number is too large, actual: %d, expected: %d", actualCount, expectedCount)
	} else {
		t.Logf("The actual and expected numbers are %d, %d", actualCount, expectedCount)
	}
}

func Test_omikuji10000times(t *testing.T) {

	resultCount := make(map[string]int)

	for i := 0; i < 10000; i++ {
		resultCount[omikuji()]++
	}

	testResultCount(t, resultCount["大吉"], 1700)
	testResultCount(t, resultCount["凶"], 3000)
	testResultCount(t, resultCount["吉"], 3600)
	testResultCount(t, resultCount["末吉"], 600)
	testResultCount(t, resultCount["末小吉"], 300)
	testResultCount(t, resultCount["半吉"], 400)
	testResultCount(t, resultCount["小吉"], 300)
}
