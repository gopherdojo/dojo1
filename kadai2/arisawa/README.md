# 課題2

## 課題2-1

io.Readerとio.Writerについて調べてみよう。

* 標準パッケージでどのように使われているか
* io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

## 課題2-2

課題1のテストを作ってみて下さい。

* テストのしやすさを考えてリファクタリングしてみる
* テストのカバレッジを取ってみる
* テーブル駆動テストを行う
* テストヘルパーを作ってみる

# 説明

## 課題2-1

（なんと書いたか忘れてしまっているが、以下のようなことを書いたと思います）

* 画像のデコード、エンコード周りで使われている。
* 構造体や変数にio.Reader/Writerインターフェースを実装することで、開発者は具体的な実装を知らなくても利用できるところが利点。
  * 例えば png.Encode や os.Stdout 

## 課題2-2

* テストのしやすさを考えてリファクタリングしてみる
  * 最初に作った時は `Do()` というメソッドがまるっと全てやる実装になっていた
  * まず、サポートしている画像フォーマットを持つ型を作って、それにフォーマットを検査するメソッドをはやした
  * `imgconv.Convert(src, dest string)` という変換だけやる関数を作ってそのテストを書いた
  * `type RecursiveConverter struct` という構造体を作り、そちらで対象ファイルの抽出、変換ファイル名の生成をするようにした
  * `imgconv.Convert()` は、内部で `t := &target{src, dest}; t.convert()` するようにし、targets 構造体を RecursiveCoverter の属性に持たせた
  * `NewRecursiveConverter` では targets 埋めまでやるようにしたので、それをテストするようにした
    * これが大きくなってしまったこれじゃない感ある
* テストのカバレッジを取ってみる
  * 変換フォーマットの組み合わせに対応していないので、そこでカバレッジ下げてる予感
```
➜  go-imgconv git:(master) ✗ go test -coverprofile=profile ./imgconv
ok      github.com/arisawa/go-imgconv/imgconv   0.323s  coverage: 82.7% of statements
```
* テーブル駆動テストを行う
  * テストはテーブル駆動テストにしている
* テストヘルパーを作ってみる
  * うっ やってない

* 拡張した点
  * 誰もほしくない `from のみ webp 変換`
  * 準標準パッケージにあったが、Encode() がなかったので、出力ができない

* 気になる点
  * private なメソッドをテストで参照しなくてもいいように書く、というのはなんとなく慣れでできている
  * テストしたい構造体の属性の値があったが、それをどうやってテストするか迷った
    * 結局 GetTargets() とか Getter生やして対応したけどいい書き方なかったかなと思っている

以下は godocdown で生成したものなので、上に課題コメント書いた。

---

# imgconv
--
    import "github.com/arisawa/go-imgconv/imgconv"


## Usage

```go
var DestFormats = Formats{"png", "jpg", "gif"}
```
DestFormats is the list of supported destination formats.

```go
var SourceFormats = Formats{"png", "jpg", "gif", "webp"}
```
SourceFormats is the list of supported source formats.

#### func  Convert

```go
func Convert(src, dest string) error
```
Convert executes image conversion a source file to the destination file.

#### type Formats

```go
type Formats []string
```

Formats is the list of registered image formats.

#### func (*Formats) Inspect

```go
func (f *Formats) Inspect(file string) bool
```
Inspect returns true value when image format is supported.

#### type RecursiveConverter

```go
type RecursiveConverter struct {
}
```

RecursiveConverter converts target images recursively.

#### func  NewRecursiveConverter

```go
func NewRecursiveConverter(in, out, srcFormat, destFormat string) (*RecursiveConverter, error)
```
NewRecursiveConverter allocates a new RecursiveConverter struct and detect
error.

#### func (*RecursiveConverter) Convert

```go
func (rc *RecursiveConverter) Convert() error
```
Convert executes image conversion for target files.

#### func (*RecursiveConverter) GetTargets

```go
func (rc *RecursiveConverter) GetTargets() []*target
```
GetTargets returns property of targets.
