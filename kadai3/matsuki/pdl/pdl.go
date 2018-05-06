package pdl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/sync/errgroup"
)

// Client structs
type Client struct {
	URL            *url.URL
	URLStr         string
	HTTPClient     *http.Client
	isParallelMode bool
	Procs          uint
	RangeSize      uint
	timeout        int
	filename       string
}

// Range struct for range access
type Range struct {
	low    uint
	high   uint
	worker uint
}

// NewClient is setter paralle downloader
func NewClient(targetURL string, isParallel bool, procs uint) (*Client, error) {
	trgURL, err := url.ParseRequestURI(targetURL)
	if err != nil {
		return &Client{}, err
	}

	if procs == uint(0) {
		cpus := runtime.NumCPU()
		max := runtime.GOMAXPROCS(cpus)
		procs = uint(max)
		warn(fmt.Sprint("set procsNum by CPU:", procs))
	}
	if isParallel && procs <= 1 {
		warn(fmt.Sprint("procsNum = 1. then exec singleDownload(): p=", isParallel, " / procsNum =", procs))
		isParallel = false
	}

	fn := getFilenameOfURL(targetURL)

	pdl := &Client{
		URL:            trgURL,
		URLStr:         targetURL,
		HTTPClient:     &http.Client{},
		Procs:          procs,
		RangeSize:      0, // set nil value
		timeout:        60,
		isParallelMode: isParallel,
		filename:       fn,
	}
	return pdl, nil
}

// Download is exec download with isParallelMode
func (p *Client) Download() error {
	if p.isParallelMode {
		if err := p.Checking(); err != nil {
			return err
		}
		return p.ParallelDownload()
	}
	return p.SingleDownload()
}

// Checking is prepare
func (p *Client) Checking() error {
	ctx, cancelAll := context.WithTimeout(context.Background(), time.Duration(p.timeout)*time.Second)
	defer cancelAll()

	size, err := p.CheckAllowRangeAccess(ctx)
	if err != nil {
		return err
	}
	p.RangeSize = uint(size)

	debug(fmt.Sprint("Cheking() for targetURL. rangeAccess is OK, size=", size, ", err=", err))
	return nil
}

// CheckAllowRangeAccess check  header(can Accept-Raneges) and return fizesize
func (p *Client) CheckAllowRangeAccess(ctx context.Context) (int64, error) {
	res, err := ctxhttp.Head(ctx, http.DefaultClient, p.URLStr)
	if err != nil {
		return 0, errors.Wrap(err, "failed to head request: "+p.URLStr)
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.Errorf("not supported range access: %s", p.URLStr)
	}
	// get of ContentLength
	if res.ContentLength <= 0 {
		return 0, errors.New("invalid content length")
	}
	return res.ContentLength, nil
}

func isFilePartsExist(filename string, r Range) (bool, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	requiredSize := int64(r.high - r.low)
	debug(fmt.Sprint("checking cache file.. have -> ", fi.Size(), " / want ->", requiredSize))
	if fi.Size() < requiredSize {
		return false, errors.New(fmt.Sprint("filesize is different. want(", requiredSize, "), but have(", fi.Size(), ")"))
	}
	return true, nil
}

// ParallelDownload is exec normal download
func (p *Client) ParallelDownload() error {
	var eg errgroup.Group
	for i := uint(0); i < p.Procs; i++ {
		pnum := i // redefine in this loop
		eg.Go(func() error {
			r := p.makeRange(pnum)
			filename := fmt.Sprintf("%d_%d_%s", p.Procs, r.worker, p.filename)

			// check cache file
			if ok, _ := isFilePartsExist(filename, r); ok {
				warn(fmt.Sprint("use cache file of", filename))
				return nil
			}
			err := os.Remove(filename)
			if err != nil {
				debug(fmt.Sprint("delete error at ", filename, " of: ", err))
			}

			debug(fmt.Sprint("range access start", filename, "with range of", r))
			res, err := p.GetResponse(r) // range access
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to split get requests: %d", r.worker))
			}
			defer res.Body.Close()

			output, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				return errors.Wrapf(err, "failed to create in %s", filename)
			}
			defer output.Close()

			if _, err := io.Copy(output, res.Body); err != nil {
				return errors.Wrapf(err, "failed to copy in %s", filename)
			}
			debug(fmt.Sprint("downloaded of ", filename))
			return nil
		})
	} // End of for
	if err := eg.Wait(); err != nil {
		errorlog(fmt.Sprint("range access err of", err))
	}

	finalFile := p.filename
	if err := os.Remove(finalFile); err != nil {
		debug(fmt.Sprint("delete err at ", finalFile, " of:", err))
	}

	debug(fmt.Sprint("create finalFile of ", finalFile))
	err := getherSplitedFiles(finalFile, p.Procs)
	if err != nil {
		return err
	}
	return nil
}

// GetResponse is Range access
func (p *Client) GetResponse(r Range) (*http.Response, error) {
	req, err := http.NewRequest("GET", p.URLStr, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to split NewRequest for get: %d", r.worker))
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", r.low, r.high))
	return http.DefaultClient.Do(req)
}

func getherSplitedFiles(baseFilename string, procs uint) error {
	fh, err := os.Create(baseFilename)
	if err != nil {
		return errors.Wrapf(err, "failed to create a file in download location of %s", baseFilename)
	}
	defer fh.Close()

	var f string
	for i := uint(0); i < procs; i++ {
		f = fmt.Sprintf("%d_%d_%s", procs, i, baseFilename)
		subfp, err := os.Open(f)
		if err != nil {
			return errors.Wrapf(err, "getherSplitFiles() with %s", f)
		}

		io.Copy(fh, subfp)
		subfp.Close()
	}
	return nil
}

// SingleDownload is exec normal download
func (p *Client) SingleDownload() error {
	debug(fmt.Sprint("SingleDownload() with: ", p.URLStr))
	res, err := p.HTTPClient.Get(p.URLStr)
	// res, err := http.Get(p.URLStr)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("http status err =" + res.Status)
	}
	warn(fmt.Sprint("status=", res.Status))
	defer res.Body.Close()

	err = os.Remove(p.filename)
	if err != nil {
		debug(fmt.Sprint("failed to delete :", p.filename, err))
	}
	fp, err := os.OpenFile(p.filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrapf(err, "save() failed with ", p.filename)
	}
	defer fp.Close()

	if _, err := io.Copy(fp, res.Body); err != nil {
		return errors.Wrapf(err, "failed to copy in %s", p.filename)
	}
	return nil
}

// Filename return c.filename
func (p *Client) Filename() string {
	return p.filename
}

// IsParallelMode return c.isParallelMode
func (p *Client) IsParallelMode() bool {
	return p.isParallelMode
}

func (p *Client) makeRange(i uint) Range {
	split := p.RangeSize / p.Procs
	low := split * i
	high := low + split - 1
	if i == p.Procs-1 {
		high = p.RangeSize
	}
	return Range{
		low:    low,
		high:   high,
		worker: i,
	}
}
func getFilenameOfURL(targetURL string) string {
	_, filename := path.Split(targetURL)
	if filename == "" {
		filename = "downloaded_file"
	}
	return filename
}
