# imgconv

指定したファイルを再帰的に調査し、画像を変換します。

実行がとても遅いので、実際のプログラム上で利用するのは避けてください。

[![GoDoc](https://godoc.org/github.com/srttk/imgconv/converter?status.svg)](https://godoc.org/github.com/srttk/imgconv/converter)

## Installation

`$ go get github.com/srttk/imgconv`

## Usage

`$ imgconv -src=jpg -out=png IMG_DIR`

### src引数

`src` 属性は `png, jpg` などの変換したい画像の拡張子です。

この属性に指定したファイルタイプの画像のみ変換します。

対応している拡張子は `png, jpg, jpeg, gif` です。

`-s` ショートハンドがあります

### out属性

`out` 属性もおなじく、ファイルタイプを指定する拡張子を入力します。

この属性で指定した拡張子に画像を変換します。

対応している拡張子は `png, jpg, jpeg, gif` です。

`-o` ショートハンドがあります

## Author

Yuya Okumur (srttk)
