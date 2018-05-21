package iria

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/sync/errgroup"
)

//ParallelDownloader is interface
type ParallelDownloader interface {
	Execute() error
	GetContentLength() error
	Head(ctx context.Context) (*http.Response, error)
	SplitDownload(part int, rangeString string) error
	MargeChunk() error
	CleanUp() error
}

//Downloader implemants ParallelDownloader
type Downloader struct {
	URL      string //取得対象URL
	SplitNum int    //ダウンロード分割数
}

//ダウンロード用一時ファイル part1~part{splitNum}
const tmpFile = "part"

//Execute はDownloaderメイン処理
func (d *Downloader) Execute() error {
	eg, ctx := errgroup.WithContext(context.Background())
	//取得対象リソースサイズ取得
	contentLength, err := d.GetContentLength()
	if err != nil {
		return err
	}
	//gorutineで分割ダウンロード
	for i, v := range getByteRange(contentLength, d.SplitNum) {
		part := i + 1
		rangeString := v
		log.Printf("splitDownload part%v start %v\n", i+1, v)
		//goルーチンで動かす関数や処理はforループが回りきってから動き始める(引数も回りきった後の状態)ので
		//goルーチン内でAdd(1)するとWaitされない場合がある
		eg.Go(func() error {
			return d.SplitDownload(ctx, part, rangeString)
		})
	}
	defer d.CleanUp()
	//分割ダウンロードが終わるまでブロック
	if err := eg.Wait(); err != nil {
		return err
	}
	//分割ダウンロードしたファイル合体
	margeFile, err := os.Create(filepath.Base(d.URL))
	if err != nil {
		return err
	}
	defer margeFile.Close()

	return d.MargeChunk(margeFile)
}

//GetContentLength は取得対象リソースのサイズを取得する
func (d *Downloader) GetContentLength() (int64, error) {
	//ファイルのサイズを取得
	//Content-TypeとContent-Length
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := d.Head(ctx)
	if err != nil {
		return 0, err
	}
	if http.StatusOK != res.StatusCode {
		return 0, fmt.Errorf("URL:%v Status:%v", d.URL, res.StatusCode)
	}
	if "bytes" != res.Header.Get("Accept-Ranges") {
		return 0, fmt.Errorf("目的のリソースがrange request未対応でした:%v", d.URL)
	}
	return res.ContentLength, nil
}

//Head はテストで差し替えたいのでメソッド化
func (d *Downloader) Head(ctx context.Context) (*http.Response, error) {
	return ctxhttp.Head(ctx, http.DefaultClient, d.URL)
}

//SplitDownload gorutineで並列ダウンロード
func (d *Downloader) SplitDownload(ctx context.Context, part int, rangeString string) error {
	//ファイル作成
	file, err := os.Create(fmt.Sprintf("part%v", part))
	if err != nil {
		return err
	}
	defer file.Close()
	//部分ダウンロードして外部ファイルに保存
	return partialRequest(d.URL, part, rangeString, file)
}

//分割ダウンロード
func partialRequest(url string, part int, rangeString string, w io.Writer) error {
	//リクエスト作成
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range",
		fmt.Sprintf("bytes=%v", rangeString))

	//デバッグ用リクエストヘッダ出力
	dump, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		return err
	}
	fmt.Printf("%s", dump)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("http.DefaultClient.Do(req) err:%v", err.Error())
		return err
	}
	//デバッグ用レスポンスヘッダ出力
	dumpResp, _ := httputil.DumpResponse(res, false)
	fmt.Println(string(dumpResp))

	if _, err := io.Copy(w, res.Body); err != nil {
		return err
	}
	log.Printf("partialRequest %v done", part)
	return nil
}

//MargeChunk 分割ダウンロードしたファイルを合体して復元する
//defer で定義した関数の返却値ってどうなる？
func (d *Downloader) MargeChunk(w io.Writer) error {
	for i := 0; i < d.SplitNum; i++ {
		file, err := os.Open(fmt.Sprintf("%v%v", tmpFile, i+1))
		if err != nil {
			return err
		}
		//ファイルに追記
		if _, err = io.Copy(w, file); err != nil {
			return err
		}
		if err = file.Close(); err != nil {
			return err
		}
	}
	return nil
}

//CleanUp はダウンロード用に作成した一時ファイルがあれば削除します
func (d *Downloader) CleanUp() error {
	for i := 0; i < d.SplitNum; i++ {
		target := fmt.Sprintf("%v%v", tmpFile, i+1)
		if !exists(target) {
			continue
		}
		if err := os.Remove(target); err != nil {
			return err
		}
	}
	return nil
}
