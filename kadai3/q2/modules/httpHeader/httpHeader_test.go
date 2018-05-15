package httpHeader

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"."
)

var sampleHandler = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello HTTP Test")
})

func TestLength(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	length, err := httpHeader.GetLength(ts.URL)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	// bodyのバイト数である 15 を返すこと
	if length != 15 {
		t.Fatalf("failed test %#v", length)
	}
}
