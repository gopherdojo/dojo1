package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopherdojo/dojo1/kadai4/matsuki/lib"
)

type MockDrawer struct {
}

// DrawFortuneSlip is Mock
func (m MockDrawer) DrawFortuneSlip() (string, bool) {
	return "大吉", false
}

// TestMain called first
func TestMain(t *testing.T) {

	beforeTest()

	// test here
	// TestIndex(t)

}

func beforeTest() {
	lib.SetDrawer(MockDrawer{})
}

func TestIndex(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(lib.Index))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}

	//レスポンスBODY取得
	body, _ := ioutil.ReadAll(res.Body)
	have := string(body)
	log.Print(have)

	fortuneRes := lib.FortuneResult{}
	if err := json.Unmarshal([]byte(have), &fortuneRes); err != nil {
		t.Errorf("response body cannot unmarshal. err=%v, body=%v", err, have)
	}

	// now := time.Now()
	// today := now.Format("2006-01-02")
	// want := fmt.Sprint(`{"today":"`, today, `","fortune":"大吉","is_newyear_mode":false}`)
	//
	// if have != want {
	// 	t.Errorf("response body error, want %v but have %v", want, have)
	// }
}

func TestTwice(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(lib.Twice))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}

	//レスポンスBODY取得 & JSON復元できるか
	body, _ := ioutil.ReadAll(res.Body)
	have := string(body)
	log.Print(have)

	fortuneRess := []lib.FortuneResult{}
	if err := json.Unmarshal([]byte(have), &fortuneRess); err != nil {
		t.Errorf("response body cannot unmarshal. err=%v, body=%v", err, have)
	}
}
