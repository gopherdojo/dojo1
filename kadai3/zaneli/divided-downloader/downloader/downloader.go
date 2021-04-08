package downloader

import (
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/sync/errgroup"
	"gopkg.in/cheggaaa/pb.v1"
)

type httpClient interface {
	Head(url string) (*http.Response, error)
	Do(req *http.Request) (*http.Response, error)
}

var client httpClient

func init() {
	client = http.DefaultClient
}

// Run download url resource.
func Run(url string, path string, parallel int) error {
	if parallel < 1 {
		return errors.New("並列度には1以上の値を指定してください")
	}

	length, err := getContentLength(url)
	if err != nil {
		return err
	}

	ps := make([]PartialContent, parallel)
	pbs := make([]*pb.ProgressBar, parallel)
	defer clear(ps)

	span := int(math.Ceil(float64(length) / float64(parallel)))
	for i := 0; i < parallel; i++ {
		from := i * span
		to := (i + 1) * span
		if to > length {
			to = length
		}
		pb := pb.New(to - from).Prefix(fmt.Sprintf("%s.%d", path, i))
		ps[i] = PartialContent{from: from, to: to - 1, index: i, path: path, pb: pb}
		pbs[i] = pb
	}

	pool, err := pb.StartPool(pbs...)
	if err != nil {
		return err
	}
	defer pool.Stop()

	if err = download(url, path, ps); err != nil {
		return err
	}

	return join(path, ps)
}

func download(url, path string, ps []PartialContent) error {
	eg := errgroup.Group{}

	for _, p := range ps {
		downloadParallel(url, p, &eg)
	}

	return eg.Wait()
}

func downloadParallel(url string, p PartialContent, eg *errgroup.Group) {
	eg.Go(func() error {
		p.pb.Start()

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", p.from, p.to))
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		if res == nil {
			return errors.New("response is nil")
		}
		if res.StatusCode != 206 {
			return fmt.Errorf("unexpected status code. %d", res.StatusCode)
		}

		body := res.Body
		if body == nil {
			return errors.New("response body is nil")
		}
		defer body.Close()

		file, err := os.Create(p.filePath())
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err = io.Copy(io.MultiWriter(file, p.pb), body); err != nil {
			return err
		}

		p.pb.Finish()
		return nil
	})
}

func getContentLength(url string) (int, error) {
	res, err := client.Head(url)
	if err != nil {
		return 0, err
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.New("分割ダウンロードに対応していません")
	}
	size, err := strconv.Atoi(res.Header.Get("Content-Length"))
	if err != nil {
		return 0, err
	}
	return size, nil
}

func join(path string, ps []PartialContent) error {
	files := make([]io.Reader, len(ps))

	for i, p := range ps {
		file, err := os.Open(p.filePath())
		if err != nil {
			return err
		}
		defer file.Close()
		files[i] = file
	}

	reader := io.MultiReader(files...)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = io.Copy(file, reader); err != nil {
		return err
	}
	return nil
}

func clear(ps []PartialContent) error {
	var errs []error
	for _, p := range ps {
		if err := os.Remove(p.filePath()); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) >= 1 {
		return errs[0] // TODO: いい感じに全ての errors を返したい…
	}
	return nil
}
