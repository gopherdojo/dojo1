package pdl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/matsu0228/go_sandbox/04_parallelDownloder/pdl"
)

func getSimpleTestServer() *httptest.Server {
	var sampleHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	})
	ts := httptest.NewServer(sampleHandler)
	return ts
}

// TODO: httptest cant support range access (reposense fullsize data)
func getAllowRangeaccessServer(contentLength uint, bodyText string) *httptest.Server {
	type Response struct {
		path, query, contenttype, body string
	}
	response := &Response{
		path:        "test/url",   //"/v1/media/popular",
		contenttype: "text/plain", //application/json",
		body:        bodyText,
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", response.contenttype)
		if contentLength > 0 {
			w.Header().Set("Content-Length", fmt.Sprint(contentLength))
		}
		w.Header().Add("Accept-Ranges", "bytes")
		io.WriteString(w, response.body)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	return server
}

func TestNewClient(t *testing.T) {
	testDataOK := []struct {
		testIndex   string
		trgURL      string
		isParallel  bool
		procs       uint
		resParallel bool
	}{
		// {"parallel0", "https://golang.org/", true, 0, true}, //TODO: dependecy for os.Num
		{"paralles10", "https://golang.org/", true, 10, true},
		{"single", "https://golang.org/", false, 1, false},
		{"paralles1ThenSingle", "https://golang.org/", true, 1, false},
	}
	for _, d := range testDataOK {
		t.Run(d.testIndex, func(t *testing.T) {
			p, err := pdl.NewClient(d.trgURL, d.isParallel, d.procs)
			if err != nil {
				t.Errorf("NewClient() err= %v with url: %s, isParallel: %v, procs: %v", err, d.trgURL, d.isParallel, d.procs)
			}
			isParallelMode := p.IsParallelMode()
			if isParallelMode != d.resParallel {
				t.Errorf("NewClient() parallel mode is different. isParallel: %v, procs: %v, pdl: %v", d.isParallel, d.procs, p)
			}
		})
	}
}

func TestSingleDownload(t *testing.T) {
	ts := getSimpleTestServer()
	defer ts.Close()

	c, err := pdl.NewClient(ts.URL, false, 1)
	if err != nil {
		t.Errorf("NewClient() err= %v with url: %s", err, ts.URL)
	}
	err = c.SingleDownload()
	if err != nil {
		t.Errorf("Download() err= %v", err)
	}

	// TODO: check file body
}

func TestCheckAllowRangeAccess(t *testing.T) {
	testData := []struct {
		testName       string
		testServerType string
		contentLength  uint
		body           string
		hasErr         bool
	}{
		{"rangeAccessServer", "range", 1000, "this is rangeAccess sercer.", false},
		{"simpleServer", "simple", 0, "this is simple server", true},
	}
	for _, d := range testData {
		var ts *httptest.Server
		t.Run(d.testName, func(t *testing.T) {
			if d.testServerType == "simple" {
				ts = getSimpleTestServer()
				defer ts.Close()
			} else if d.testServerType == "range" {
				ts = getAllowRangeaccessServer(d.contentLength, d.body)
				defer ts.Close()
			} else {
				panic(fmt.Sprintf("test data is invalid: %v", d))
			}

			// テスト(=CheckAllowRangeAccess)
			ctx, cancelAll := context.WithTimeout(context.Background(), (60 * time.Second))
			defer cancelAll()
			p, err := pdl.NewClient(ts.URL, true, 2)
			if err != nil {
				panic(fmt.Sprintf("NewClient() err= %v with url: %s", err, ts.URL))
			}
			size, err := p.CheckAllowRangeAccess(ctx)

			if d.hasErr { // errを返すパターン =ServerがRangeAccessを許容していない場合
				if err == nil {
					t.Errorf("Checking() didnt return err of %v / with %v", err, p)
				}
			} else { // errを返さないパターン
				if err != nil {
					t.Errorf("Checking() return err of %v / with %v", err, p)
				}
				if size != int64(d.contentLength) {
					t.Errorf("Checking() return size of %v, but server's header is contentLength= %v", size, d.contentLength)
				}
			}

		}) // End of subtest
	}
}

func TestParallelDownload(t *testing.T) {
	bText := `1. this is test server. this is test data
----------------
----------------
2.this is test server. this is test data
----------------
----------------
3.this is test server. this is test data
----------------
----------------
`
	ts := getAllowRangeaccessServer(0, bText)
	defer ts.Close()

	c, err := pdl.NewClient(ts.URL, true, 2)
	if err != nil {
		t.Errorf("NewClient() err= %v with url: %s", err, ts.URL)
	}

	// // debug
	// resp, err := http.Head(ts.URL)
	// log.Print(resp)

	err = c.ParallelDownload()
	if err != nil {
		t.Errorf("ParallelDownload() err= %v", err)
	}
	// TODO: check res.body
	// buf, err := ioutil.ReadFile(c.Filename())
}
