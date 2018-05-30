package lib

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var drawer Drawer

// Drawer は、おみくじを引くinterface
type Drawer interface {
	DrawFortuneSlip() (string, bool)
}

// NewServer is constoructor
//  Drawer interfaceを使うことでunitTestしやすくする
func NewServer(d Drawer) http.Server {
	SetDrawer(d) // パッケージ内部グローバル変数にセット

	m := http.NewServeMux()
	// Routing
	m.Handle("/twice", http.HandlerFunc(Twice))
	m.Handle("/", http.HandlerFunc(Index))

	server := http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	log.Printf("[INFO] server start")
	return server
}

// SetDrawer is setter
func SetDrawer(d Drawer) {
	drawer = d
}

// Index では、おみくじを引き、結果をJSONで返却
func Index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		log.Printf("[INFO] index() with path = %v", req.URL.Path)
		http.NotFound(w, req)
		return
	}

	fortuneRes, isNewyearDay := drawer.DrawFortuneSlip()
	log.Printf("[INFO] index() with fourtune=%v, isNewyearDay=%v", fortuneRes, isNewyearDay)
	// w.Write([]byte("Hello world = " + fortuneRes))
	res := FortuneResult{
		Today:         jsonTime{time.Now()},
		Fortune:       fortuneRes,
		IsNewyearMode: isNewyearDay,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

// Twice では、おみくじを２回連続で引き、JSONで返却
func Twice(w http.ResponseWriter, req *http.Request) {
	fFirst, isNdFist := drawer.DrawFortuneSlip()
	fSecond, isNdSecond := drawer.DrawFortuneSlip()

	res := []FortuneResult{
		FortuneResult{
			Today:         jsonTime{time.Now()},
			Fortune:       fFirst,
			IsNewyearMode: isNdFist,
		},
		FortuneResult{
			Today:         jsonTime{time.Now()},
			Fortune:       fSecond,
			IsNewyearMode: isNdSecond,
		},
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
