package downloader

import (
	"fmt"
	"os"
	"strings"
	"io"
	"runtime"

	"golang.org/x/sync/errgroup"
	"github.com/pkg/errors"

	"github.com/gopherdojo/dojo1/kadai3/srttk/multidown/option"
)

type Downloader struct {
	OutputDir   string
	FileName    string
	MaxProcess  uint
	Workers     []*worker
}

func NewDownloader(cliOption option.CliOption, size uint, url string) (*Downloader, error) {
	d := new(Downloader)
	validatedOutputDir, err := getValidatedOutputDir(cliOption.OutputDir)
	if err != nil {
		return nil, err
	}
	d.OutputDir = validatedOutputDir
	d.FileName = getFileName(url)
	d.MaxProcess = uint(runtime.NumCPU())
	split := size / d.MaxProcess
	for i := uint(0); i < d.MaxProcess; i++ {
		w, err := NewWorker(d, size, i, split, url)
		if err != nil {
			return nil, errors.Wrap(err, "initialize worker error")
		}
		d.Workers = append(d.Workers, w)
	}
	fmt.Fprintf(os.Stdout, "Download start from %s\n", url)
	return d, nil
}

func getValidatedOutputDir(outputDir string) (string, error) {
	_, err := os.Open(outputDir)
	if err != nil {
		return "", err
	}
	return outputDir, nil
}

func getFileName(resourceUrl string) string {
	token := strings.Split(resourceUrl, "/")
	filename := token[len(token)-1]
	return filename
}

func (d *Downloader) Download() error {
	grp := errgroup.Group{}

	for _, worker := range d.Workers {
		w := worker
		grp.Go(func() error {
			return w.Request()
		})
	}
	// wait for Assignment method
	if err := grp.Wait(); err != nil {
		return err
	}
	return nil
}

func (d *Downloader) Merge() error {
	outputFilePath := fmt.Sprintf("%s/%s", d.OutputDir, d.FileName)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return errors.Wrap(err, "failed to create merge file")
	}
	defer outputFile.Close()
	for i := uint(0); i < d.MaxProcess; i++ {
		partFilePath := fmt.Sprintf("%s.%d", outputFilePath, i)
		partFile, err := os.Open(partFilePath)
		if err != nil {
			return errors.Wrap(err, "failed to open part file")
		}
		io.Copy(outputFile, partFile)
		partFile.Close()
		if err := os.Remove(partFilePath); err != nil {
			return errors.Wrap(err, "failed to remove a file")
		}
	}
	return nil
}
