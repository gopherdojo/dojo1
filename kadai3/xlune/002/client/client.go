package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	// UnitSize ダウンロードサイズ最小値
	UnitSize = 1024
)

// TargetInfo 対象の情報
type TargetInfo struct {
	URL      string
	Name     string
	Size     int64
	IsRanges bool
}

// DownloadRange ダウンロードレンジ
type DownloadRange struct {
	Start int64
	End   int64
}

// DownloadItem ダウンロード管理アイテム
type DownloadItem struct {
	URL        string
	RangeIndex int
	Range      DownloadRange
	Body       []byte
}

// DownloadByItem ターゲット情報でダウンロードする
func DownloadByItem(item *DownloadItem) error {

	req, err := http.NewRequest("GET", item.URL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", item.Range.Start, item.Range.End))

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	item.Body = body
	return nil
}

// MakeDownloadRange 分割ダウンロードレンジリスト生成
func MakeDownloadRange(size int64, count int) ([]DownloadRange, error) {
	result := []DownloadRange{}
	if count <= 0 {
		return result, errors.New("invalid argument")
	}
	unitCount := size / UnitSize
	if size%UnitSize > 0 {
		unitCount++
	}

	fillCount := unitCount % int64(count)
	if fillCount != 0 {
		fillCount = int64(count) - fillCount
	}
	singleCount := (unitCount + fillCount) / int64(count)

	for i := 0; i < count; i++ {
		s := int64(i) * singleCount * UnitSize
		e := int64(i+1)*singleCount*UnitSize - 1
		result = append(result, DownloadRange{
			Start: s,
			End:   int64(math.Min(float64(e), float64(size))),
		})
	}

	return result, nil
}

// GetSourceURL 実態URLの取得
func GetSourceURL(url string) (string, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 301 ||
		resp.StatusCode == 302 ||
		resp.StatusCode == 303 ||
		resp.StatusCode == 307 {

		source := resp.Header["Location"]
		if len(source) > 0 {
			return GetSourceURL(source[0])
		}
		return "", errors.New("redirect location empty")
	}
	return url, nil
}

// GetInfo ファイルサイズ取得
func GetInfo(url string) (TargetInfo, error) {
	info := TargetInfo{
		URL: url,
	}
	c := &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}

	resp, err := c.Head(url)
	info.Name = path.Base(resp.Request.URL.Path)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()

	if ar, ok := resp.Header["Accept-Ranges"]; ok && strings.ToLower(ar[0]) == "bytes" {
		info.IsRanges = true
	}

	if cl, ok := resp.Header["Content-Length"]; ok {
		size, err := strconv.ParseInt(cl[0], 10, 64)
		if err != nil {
			return info, err
		}
		info.Size = size
	}

	return info, nil
}
