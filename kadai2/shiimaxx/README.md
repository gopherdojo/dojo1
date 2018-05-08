# image-convert

[![wercker status](https://app.wercker.com/status/0b0f88b8cd21f50667fc77966fbc797b/s/master "wercker status")](https://app.wercker.com/project/byKey/0b0f88b8cd21f50667fc77966fbc797b)

## Usage

```
image-convert -s <source extension> -d <destination extention> <directory>
```

```
Usage of image-convert:
  -d string
        destination extension(Short)
  -dest string
        destination extension
  -s string
        source extension(Short)
  -src string
        source extension
  -version
        print version information
```

## Author

[shiimaxx](https://github.com/shiimaxx)

---

# io.Readerとio.Writer

## 標準パッケージでどのように使われているか

### io.Reader

課題2の画像変換プログラムで利用したimage.Decodeでどのように使われているかを見てみました。
image.Decodeは引数rにio.Readerインターフェースを実装する型の変数をとります。

rはreaderインターフェースを実装するかたちに変換されたあと、sniffの中で画像形式の判別のためにマジックナンバーの読み取りを実行しています。

コードリーディングのメモ
https://gist.github.com/shiimaxx/6d5c70c98cc83d6a085c93e4f573245a


### io.Writer

こちらも課題2の画像変換プログラムで利用したjpeg.Encodeでどのように使われているかを見てみました。
jpeg.Encodeは引数wにio.Writerインターフェースを実装する型の変数をとります。

wはwriterインターフェースを実装するかたちに変換してencoder型のeがもつe.wに代入されたあと、各種書き込み処理で利用されています。

コードリーディングのメモ
https://gist.github.com/shiimaxx/2bfa07c6572fe29f0ab141563bf651d8

---

## io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

入出力を抽象的に扱うことができのが利点と考えました。
具体的には、image.Decodeはio.Readerを、jpeg.Encodeはio.Writerを引数にとりますが、これにより実際に渡される引数の型にかかわらず、関数内ではインターフェースに実装されているRead、Writeを利用すればよいことになります。
これにより、io.Reader, io.Writerを実装するos.File、http.Response.Bodyなどを扱う際に実体を意識せずにロジックを実装できます。
