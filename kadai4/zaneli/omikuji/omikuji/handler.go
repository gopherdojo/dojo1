package omikuji

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var 大吉 = "大吉"
var 中吉 = "中吉"
var 吉 = "吉"
var 小吉 = "小吉"
var 凶 = "凶"
var 結果 = []string{大吉, 中吉, 吉, 小吉, 凶}

// Result はレスポンスJSONの構造を表す。
type Result struct {
	Result string `json:"result"`
	Date   string `json:"date"`
}

// Handler はおみくじの結果を返す。
func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	v := ctx.Value(datetimeContextKey)
	if v == nil {
		v = time.Now()
	}
	date, ok := v.(time.Time)
	if !ok {
		date = time.Now()
	}

	var result string
	if is三が日(date) {
		result = 大吉
	} else {
		rand.Shuffle(len(結果), func(i, j int) {
			結果[i], 結果[j] = 結果[j], 結果[i]
		})
		result = 結果[0]
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(Result{Result: result, Date: date.Format(time.RFC3339)})
}

func is三が日(date time.Time) bool {
	_, m, d := date.Date()
	return m == time.January && (d == 1 || d == 2 || d == 3)
}
