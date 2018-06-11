# 課題2

## 課題2-1

io.Readerとio.Writerについて調べてみよう。

* 標準パッケージでどのように使われているか
* io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

## 課題2-1 回答

* IFを柔軟にするため。具体的には、`io.Reader`を引数にとる関数があるとして、その関数に渡す引数は`io`以外にも、Interfaceを満たす`os.File`など他の型も許容される
* コードでの例を示すと下記。`bufio.NewScanner()`は、引数に`io.Reader`をとるが、`os.Stdin`を渡しても動作する
  * これは、`*File`が`io.Reader`のinterfaceに対応している（`Reader()`関数を持っている）ため
```
stdin := bufio.NewScanner(os.Stdin)

// func NewScanner(r io.Reader) *Scanner
//
// os.Stdin means NewFile(uintptr(syscall.Stdin), "/dev/stdin")
// NewFile return *File
//
// func (f *File) Read(b []byte) (n int, err error)
```

## 課題2-2

課題1のテストを作ってみて下さい。

* テストのしやすさを考えてリファクタリングしてみる
* テストのカバレッジを取ってみる
* テーブル駆動テストを行う
* テストヘルパーを作ってみる


## 課題2-2 回答

* 下記の３機能に分割し、それぞれテストしやすいIFにした
  1. convimg.New(trgDir, srcFormat, destFormat)
    * 画像変換機能への入力値のチェック ＆ `type convImageCommand{}`へのセット
  2. convimg.ConvImages()
    * 指定されたディレクトリ配下を再帰的に探索し、該当の拡張子のファイルがあれば`Convert()`する
  3. convimg.Convert(srcImg)
    * 引数で渡された画像ファイルを、変換する

## conv_img

[![CircleCI](https://circleci.com/gh/matsu0228/go_sandbox.svg?style=svg)](https://circleci.com/gh/matsu0228/go_sandbox)


* sample command
  * 下記で、画像変換ができるコマンド
```
conv_img -d *your_dir* ( -f *from_format:dest_format* )
```
* option
  * d: 必須。指定したディレクトリ配下を再帰的に処理する
  * f: デフォルト=jpg:png。変換前の形式と変換後の形式を指定できる


## test

* following
```
$ go test -cover ./convimg
ok  _/home/dev/go-tutorial/dojo1/kadai2/matsuki/convimg4.069scoverage: 86.6% of statements
```




