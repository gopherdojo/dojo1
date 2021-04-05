package omikuji

import (
	"context"
	"net/http"
	"time"
)

type contextKey string

const datetimeContextKey contextKey = "datetime"

// AddCurrentDateTime は Context に現在日時を設定する。
// refer: https://gocodecloud.com/blog/2016/11/15/simple-golang-http-request-context-example/
func AddCurrentDateTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		date := time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60))
		ctx := context.WithValue(r.Context(), datetimeContextKey, date)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
