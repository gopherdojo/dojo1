package lib

import (
	"log"
	"math/rand"
	"time"
)

// Verbos は、ログ出力レベルを制御
var Verbos bool

// Fortune は、おみくじに関する型
type Fortune struct {
	fortuneList map[uint]string
	luckiest    string
	today       time.Time
}

// NewFortune は、おみくじの初期化
func NewFortune(now time.Time) Fortune {
	// おみくじのタネを初期化
	rand.Seed(time.Now().UnixNano())

	return Fortune{
		fortuneList: map[uint]string{
			0: "大吉",
			1: "中吉",
			2: "吉",
			3: "凶",
			4: "大凶",
		},
		luckiest: "大吉",
		today:    now,
	}
}

func (f Fortune) isNewyearsDay() bool {
	if Verbos {
		log.Printf("[DEBUG] today is %v %v", f.today.Day(), f.today.Month())
	}

	// Month is defined here:  https://golang.org/pkg/time/#Month
	if f.today.Month().String() == "January" {
		switch f.today.Day() {
		case 1, 2, 3:
			return true
		}
	}
	return false
}

// DrawFortuneSlip は、おみくじを引く関数
//  return ( 運勢, is正月 )
func (f Fortune) DrawFortuneSlip() (string, bool) {
	// 正月には大吉を引く
	if f.isNewyearsDay() {
		return f.luckiest, true
	}
	// 等確率でひく
	r := rand.Intn(len(f.fortuneList))
	return f.fortuneList[uint(r)], false
}
