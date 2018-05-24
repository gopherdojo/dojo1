package main

import (
	"net/http"
	"math/rand"
	"time"
	"encoding/json"
)

type LuckJson struct {
	Luck string `json:"luck"`
}

var t int64

func init() {
	t = time.Now().UnixNano()
	rand.Seed(t)
}

func getRand() int {
	return rand.Intn(6) + 1
}

func drawFortune(num int) string {
	switch num {
	case 6:
		return "大吉"
	case 5, 4:
		return "中吉"
	case 3, 2:
		return "吉"
	case 1:
		return "凶"
	default:
		return "Unknown Number"
	}
}

func isNewYear(_ int, month time.Month, day int) bool {
	if month == time.January && 1 <= day && day <= 3 {
		return true
	}
	return false
}

func getLuck(currentTime time.Time) (luck string) {
	if isNewYear(currentTime.Date()) {
		luck = "大吉"
	} else {
		random := getRand()
		luck = drawFortune(random)
	}
	return luck
}

func fortune(w http.ResponseWriter, r *http.Request) {
	luckJson := new(LuckJson)
	luckJson.Luck = getLuck(time.Now())
	json.NewEncoder(w).Encode(luckJson)
}

func main() {
	http.HandleFunc("/", fortune)
	http.ListenAndServe(":8080", nil)
}