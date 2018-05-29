package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"math/rand"
	"encoding/json"
	_ "net/http/pprof"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Responce struct {
	Omikuji    string `json:"omikuji"`
}

type Routes []Route

var routes = Routes{
	Route{
		"Getomikuji",
		"GET",
		"/api/v1/omikuji",
		GetOmikuji,
	},
}

var timeNowFunc = time.Now // 関数をグローバル変数に代入


func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}

func GetOmikuji(w http.ResponseWriter, r *http.Request){
	var responce Responce
	t := timeNowFunc()

	// 1/1 ~ 1/3は絶対大吉
	if (t.Month() == 1 && t.Day() >= 1 && t.Day() <= 3){
		responce.Omikuji = "大吉"
	} else {
		t := timeNowFunc().UnixNano()
		rand.Seed(t)
		s := rand.Intn(6)
		switch s + 1 {
		case 1:
			responce.Omikuji = "凶"
		case 2, 3:
			responce.Omikuji = "吉"
		case 4, 5:
			responce.Omikuji = "中吉"
		case 6:
			responce.Omikuji = "大吉"
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(responce); err != nil {
		panic(err)
	}

}
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}