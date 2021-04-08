package downloader

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

const targetURL = "http://example.com/dummy"

func Test_分割ダウンロード成功(t *testing.T) {
	basePath, err := createTmpPath("dummy")
	if err != nil {
		t.Fatal(err)
	}

	for _, parallel := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		path := fmt.Sprintf("%s.%d", basePath, parallel)
		defer os.Remove(path)

		body := createRandomString()
		c := newClientForTest(body)
		client = &c

		err = Run(targetURL, path, parallel)
		if err != nil {
			t.Fatal(err)
		}
		if c.doingCount != parallel {
			t.Errorf("unexpected client.Do() executing count: expected=%d, actual=%d", parallel, c.doingCount)
		}

		file, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			t.Fatal(err)
		}
		if string(content) != body {
			t.Errorf("unexpected file content: expected=%s, actual=%s", body, string(content))
		}
	}
}

func Test_不正な並列度を指定した場合エラー(t *testing.T) {
	path, err := createTmpPath("invalid-parallel")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	c := newClientForTest(createRandomString())
	client = &c

	err = Run(targetURL, path, 0)
	if err == nil || err.Error() != "並列度には1以上の値を指定してください" {
		t.Errorf("unexpected error: err=%v", err)
	}
}

func Test_AcceptRangesヘッダがない場合エラー(t *testing.T) {
	path, err := createTmpPath("invalid-header")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	c := newClientForTest(createRandomString())
	c.header.Del("Accept-Ranges")
	client = &c

	err = Run(targetURL, path, 6)
	if err == nil || err.Error() != "分割ダウンロードに対応していません" {
		t.Errorf("unexpected error: err=%v", err)
	}
}

func Test_206以外のステータスが返された場合エラー(t *testing.T) {
	path, err := createTmpPath("invalid-status")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	c := newClientForTest(createRandomString())
	c.statusCode = 200
	client = &c

	err = Run(targetURL, path, 6)
	if err == nil || err.Error() != fmt.Sprintf("unexpected status code. %d", c.statusCode) {
		t.Errorf("unexpected error: err=%v", err)
	}
}

func createTmpPath(name string) (string, error) {
	dir, err := ioutil.TempDir("", "divided-downloader")
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(filepath.Join(dir, name))
	if err != nil {
		return "", err
	}
	return path, nil
}

func createRandomString() string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, rand.Intn(10000)+1)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

type httpClientForTest struct {
	header     http.Header
	statusCode int
	body       string
	doingCount int
}

func newClientForTest(body string) httpClientForTest {
	return httpClientForTest{
		header: map[string][]string{
			"Accept-Ranges":  []string{"bytes"},
			"Content-Length": []string{strconv.Itoa(len(body))},
		},
		statusCode: 206,
		body:       body,
	}
}

func (c *httpClientForTest) Head(url string) (*http.Response, error) {
	if url != targetURL {
		return nil, fmt.Errorf("unexpected url: expected=%s, actual=%s", targetURL, url)
	}
	res := http.Response{}
	res.Header = c.header
	return &res, nil
}

var r = regexp.MustCompile(`bytes=(\d+)-(\d+)`)

func (c *httpClientForTest) Do(req *http.Request) (*http.Response, error) {
	if req.URL.String() != targetURL {
		return nil, fmt.Errorf("unexpected url: expected=%s, actual=%s", targetURL, req.URL)
	}

	ranges := r.FindAllStringSubmatch(req.Header.Get("Range"), -1)[0]
	start, err := strconv.Atoi(ranges[1])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(ranges[2])
	if err != nil {
		return nil, err
	}

	res := http.Response{}
	res.StatusCode = c.statusCode
	res.Body = ioutil.NopCloser(strings.NewReader(c.body[start : end+1]))

	c.doingCount++
	return &res, nil
}
