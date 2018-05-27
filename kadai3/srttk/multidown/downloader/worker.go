package downloader

import (
	"fmt"
	"os"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type worker struct {
	processId            uint
	bytesToStartReading  uint
	bytesToFinishReading uint
	resourceUrl          string
	partFilePath         string
}

func NewWorker(d *Downloader, size uint, i uint, split uint, url string) (*worker, error) {
	bytesToStartReading := split * i
	bytesToFinishReading := bytesToStartReading + split - 1
	partFilePath := fmt.Sprintf("%s/%s.%d", d.OutputDir, d.FileName, i)

	if isExists(partFilePath) {
		return nil, errors.New("part file is exists")
	}

	if i == d.MaxProcess-1 {
		bytesToFinishReading = size
	}

	//それぞれのワーカーについて、担当範囲と番号を返す
	w := &worker{
		processId: i,
		bytesToStartReading: bytesToStartReading,
		bytesToFinishReading: bytesToFinishReading,
		resourceUrl: url,
		partFilePath: partFilePath,
	}
	return w, nil
}

func isExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func (w *worker) Request() error {
	res, err := w.MakeResponse()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to split get requests: %d", w.processId))
	}
	defer res.Body.Close()

	output, err := os.Create(w.partFilePath)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to create file %s", w.partFilePath))
	}
	defer output.Close()
	//内容をコピーする
	io.Copy(output, res.Body)
	return nil
}

func (w *worker) MakeResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", w.resourceUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to split NewRequest for get: %d", w.processId))
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", w.bytesToStartReading, w.bytesToFinishReading))
	return http.DefaultClient.Do(req)
}