package client_test

import (
	"testing"

	"github.com/xlune/dojo1/kadai3/xlune/002/client"
)

const (
	targetUrl            = "https://dl.google.com/go/go1.10.2.src.tar.gz"
	targetLightWeightUrl = "https://getbootstrap.com/docs/4.0/getting-started/download/"
)

func TestDownloadByTargetInfo(t *testing.T) {
	url, err := client.GetSourceURL(targetLightWeightUrl)
	if err != nil {
		t.Fatal("Get Header failed")
	}
	info, err := client.GetInfo(url)
	if err != nil {
		t.Fatal("Get TargetInfo failed")
	}

	item := &client.DownloadItem{
		URL:        url,
		RangeIndex: 0,
		Range: client.DownloadRange{
			Start: 0,
			End:   info.Size,
		},
	}
	err = client.DownloadByItem(item)
	if err != nil {
		t.Fatal("Download failed")
	}
}

var downloadRangeSample = []struct {
	size    int64
	count   int
	ranges  []client.DownloadRange
	isError bool
}{
	{1234567890, 1, []client.DownloadRange{
		client.DownloadRange{Start: 0, End: 1234567890},
	}, false},
	{1234567890, 0, []client.DownloadRange{
		client.DownloadRange{Start: 0, End: 0},
	}, true},
	{1234567890, 3, []client.DownloadRange{
		client.DownloadRange{Start: 0, End: 411523071},
		client.DownloadRange{Start: 411523072, End: 823046143},
		client.DownloadRange{Start: 823046144, End: 1234567890},
	}, false},
}

func TestMakeDownloadRange(t *testing.T) {
	for _, rs := range downloadRangeSample {
		dr, err := client.MakeDownloadRange(rs.size, rs.count)
		if (err != nil) != rs.isError {
			t.Fatal("error result not match")
		}
		if !rs.isError {
			for index, r := range rs.ranges {
				if r != dr[index] {
					t.Fatal("range not match")
				}
			}
		}
	}
}

func TestGetSourceURL(t *testing.T) {
	_, err := client.GetSourceURL(targetUrl)
	if err != nil {
		t.Fatal("Get Header failed")
	}

}

func TestGetInfo(t *testing.T) {
	url, err := client.GetSourceURL(targetUrl)
	if err != nil {
		t.Fatal("Get Header failed")
	}

	_, err = client.GetInfo(url)
	if err != nil {
		t.Fatal("Get TargetInfo failed")
	}
}
