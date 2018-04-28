package convertor

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

//Usage of goConvImgExtention
var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, SupportDescription())
	fmt.Fprintln(os.Stderr, "<コマンド実行例>")
	fmt.Fprintln(os.Stderr, "------------------------------------------------")
	fmt.Fprintln(os.Stderr, "$goConvImgExtention -f jpg -t png {targetPath}")
	fmt.Fprintln(os.Stderr, "------------------------------------------------")
}

//Convertor struct
type Convertor struct {
	From       string
	To         string
	TargetPath string
}

//New コンバーターインスタンス作成
func New(args []string) (*Convertor, error) {
	/*
		引数とフラグの処理
	*/
	flags := flag.NewFlagSet("convertor", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	from := flags.String("f", "jpg", "convert from")
	to := flags.String("t", "png", "convert to")
	flags.Parse(args[1:])

	if len(flags.Args()) != 1 {
		return nil, errors.New("変換対象とするディレクトリを１つ指定してください")
	}
	//引数のディレクトリの存在チェック
	_, err := os.Stat(flags.Args()[0])
	if err != nil {
		return nil, err
	}
	//サポート対象外の拡張子が指定された場合nilを返す
	if ImageExtention[*from] && ImageExtention[*to] {
		return &Convertor{
			From:       *from,
			To:         *to,
			TargetPath: flags.Args()[0],
		}, nil
	}
	return nil, errors.New("サポート対象外の画像形式が指定されています。")
}

//Convert は引数で指定されたディレクトリ配下のファイルの拡張子を再帰的にFrom->Toに変換します。
func (c *Convertor) Convert() error {
	if err := filepath.Walk(c.TargetPath, c.process); err != nil {
		return errors.Wrapf(err, "%v配下の変換時に問題が起きました\n%v\n", c.TargetPath)
	}
	return nil
}

//process はWalkから呼ばれるコールバック関数
func (c *Convertor) process(path string, info os.FileInfo, err error) error {
	//ディレクトリの場合スルー
	if err != nil || info.IsDir() {
		return nil
	}
	//変換対象の拡張子でなければスルー
	ext := strings.TrimLeft(filepath.Ext(path), ".")
	if ext != c.From {
		return nil
	}
	//変換元画像処理
	imgFileFrom, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "%v could not open:%v\n", path, err.Error())
	}
	defer imgFileFrom.Close()
	imgFrom, _, err := image.Decode(imgFileFrom)
	if err != nil {
		return errors.Wrapf(err, "%v could not decode:%v\n", path, err.Error())
	}

	//変換先のファイル作成
	rename := strings.TrimRight(path, c.From) + c.To
	imgFileTo, err := os.Create(rename)
	if err != nil {
		return errors.Wrapf(err, "%v could not create:%v\n", rename, err.Error())
	}
	defer imgFileTo.Close()

	//形式に応じてエンコード
	switch ext {
	case "jpeg", "jpg":
		err = jpeg.Encode(imgFileTo, imgFrom, nil)
	case "gif":
		err = gif.Encode(imgFileTo, imgFrom, nil)
	case "png":
		err = png.Encode(imgFileTo, imgFrom)
	}
	log.Printf("convert %v -> %v\n", path, rename)
	return err
}
