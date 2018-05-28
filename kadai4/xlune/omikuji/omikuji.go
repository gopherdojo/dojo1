package omikuji

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Fortune おみくじ結果タイプ
type Fortune int

// JSONTime JSON出力時間
type JSONTime struct {
	time.Time
}

const (
	// Kyo 凶
	Kyo Fortune = iota
	// Kichi 吉
	Kichi
	// Chukichi 中吉
	Chukichi
	// Daikichi 大吉
	Daikichi
)

// Result 返却データ
type Result struct {
	Label string   `json:"label"`
	Type  Fortune  `json:"type"`
	Date  JSONTime `json:"date"`
}

// MarshalJSON JSON文字列生成
func (t JSONTime) MarshalJSON() ([]byte, error) {
	timeStr := fmt.Sprintf("\"%s\"", t.Time.Format("2006-01-02 15:04:05"))
	return []byte(timeStr), nil
}

// UnmarshalJSON JSON文字列生成パース
func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	timeObj, err := strToTime(strings.Trim(string(data), "\""))
	if err != nil {
		return err
	}
	t.Time = timeObj
	return nil
}

// Handler ハンドラ実装
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	dateStr, ok := r.URL.Query()["date"]
	dateTime := time.Now()

	if ok && len(dateStr) > 0 {
		t, err := strToTime(dateStr[0])
		if err == nil {
			dateTime = t
		}
	}

	fortune := getFortune(dateTime)
	result := Result{
		Type:  fortune,
		Label: getLabel(fortune),
		Date:  JSONTime{dateTime},
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Println("Error:", err)
	}
}

func getFortune(t time.Time) Fortune {
	// 1/1 - 1/3 は大吉
	if t.Month() == time.January && t.Day() <= 3 {
		return Daikichi
	}
	// Seed指定でランダム選択
	rand.Seed(t.UnixNano())
	n := rand.Intn(4)
	return Fortune(n)
}

func getLabel(f Fortune) string {
	switch f {
	case Kichi:
		return "吉"
	case Chukichi:
		return "中吉"
	case Daikichi:
		return "大吉"
	default:
		return "凶"
	}
}

func strToTime(str string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		t, err = time.Parse("2006/01/02 15:04:05", str)
		if err != nil {
			return time.Time{}, err
		}
	}
	return t, nil
}
