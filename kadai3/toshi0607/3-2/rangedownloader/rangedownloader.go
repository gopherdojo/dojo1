package rangedownloader

import (
	"net/http"
	"strconv"
	"fmt"
	"os"
	"path"
	"io"
	"sync"
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/pkg/errors"
	"github.com/jessevdk/go-flags"
	"strings"
)

const tempDir = "dlTmp"

var wg sync.WaitGroup

type Downloader struct {
	Argv  []string
	procs int
	url   string
	name  string
}

type cliOptions struct {
	Name  string `short:"n" long:"name" description:"output file name with extension. if not provided, rangedownloader will guess a file name based on URL"`
	Procs int    `short:"p" long:"procs" description:"number of parallel" default:"0"`
	Args struct {
		URL string
	} `positional-args:"yes"`
}

func New() *Downloader {
	return &Downloader{Argv: os.Args[1:]}
}

func (d *Downloader) Run() int {
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		fmt.Println(err)
		return 1
	}

	if err := d.parseCommandLine(); err != nil {
		fmt.Println(err)
		return 1
	}

	len, err := d.getContentLength()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	subFileLen := len / d.procs
	remain := len % d.procs

	eg, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < d.procs; i++ {
		i := i

		from := subFileLen * i
		to := subFileLen * (i + 1)

		if i == d.procs-1 {
			to += remain
		}

		eg.Go(func() error {
			return d.rangeRequest(ctx, from, to, i)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
		return 1
	}

	if err := d.createFile(); err != nil {
		fmt.Println(err)
		return 1
	}

	if err := os.RemoveAll(tempDir); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func (d *Downloader) parseCommandLine() error {
	opts := &cliOptions{}
	p := flags.NewParser(opts, flags.HelpFlag)
	_, err := p.ParseArgs(d.Argv)
	if err != nil {
		return err
	}

	d.url = opts.Args.URL
	if opts.Name == "" {
		if name := d.guessFileName(); name == "" {
			return errors.Wrap(err, "please provide output file name")
		} else {
			d.name = name
		}
	} else {
		d.name = opts.Name
	}
	d.procs = opts.Procs
	return nil
}

func (d *Downloader) guessFileName() string {
	s := strings.Split(d.url, "/")
	return s[len(s)-1]
}

func (d *Downloader) getContentLength() (int, error) {
	res, err := http.Head(d.url)
	if err != nil {
		return 0, err
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.New("this site doesn't support a range request")
	}
	len, err := strconv.Atoi(res.Header.Get("Content-Length"))
	fmt.Printf("total length: %d bytes\n", len)
	if err != nil {
		return 0, err
	}
	return len, nil
}

func (d *Downloader) rangeRequest(ctx context.Context, from int, to int, i int) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", d.url, nil)
	if err != nil {
		return err
	}

	rangeHeader := "bytes=" + strconv.Itoa(from) + "-" + strconv.Itoa(to-1)
	req.Header.Add("Range", rangeHeader)
	// errgroup.WithContext wraps context by calling context.WithCancel
	// cf. https://github.com/golang/sync/blob/master/errgroup/errgroup.go#L34
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("Range: %v, %v bytes\n", rangeHeader, resp.ContentLength)
	defer resp.Body.Close()

	file, err := os.OpenFile(path.Join(tempDir, fmt.Sprint(i)), os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, resp.Body)

	return nil
}

func (d *Downloader) createFile() error {
	file, err := os.Create(d.name)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; i < d.procs; i++ {
		subFile, err := os.Open(path.Join(tempDir, fmt.Sprint(i)))
		if err != nil {
			return err
		}
		io.Copy(file, subFile)

		subFile.Close()
	}
	return nil
}
