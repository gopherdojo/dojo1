package omikuji

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var nowFunc = time.Now

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Response is the type of omikuji result
type Response struct {
	Result string `json:"result"`
	Date   string `json:"date"`
}

// HTTPHandler writes omikuji into HTTP response
func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := Response{
		Result: omikuji(),
		Date:   nowFunc().String(),
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("Error in JSON Encoder: ", err)
	}
}

func omikuji() string {

	now := nowFunc()

	if 1 <= now.YearDay() && now.YearDay() <= 3 {
		return "大吉！"
	}

	i := rand.Intn(100)
	switch {
	case i < 17:
		return "大吉"
	case i < 17+30:
		return "凶"
	case i < 17+30+36:
		return "吉"
	case i < 17+30+36+6:
		return "末吉"
	case i < 17+30+36+6+3:
		return "末小吉"
	case i < 17+30+36+6+3+4:
		return "半吉"
	default:
		return "小吉"
	}
}
