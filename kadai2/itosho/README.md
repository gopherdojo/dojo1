# gonverter
gonverter is a CLI tool for converting image files.

## Usage
```
$ gonverter [-f from extension] [-t to extension] [directory]
```

## License
- [NYSL](http://www.kmonos.net/nysl/)

## Homework

### io.Readerとio.Writerについて調べてみよう

#### 標準パッケージでどのように使われているか
osパッケージのFile型の構造体やbufioパッケージのReader型とWriter型の構造体などがそれぞれのinterfaceを実装している。

#### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
Read / Writeメソッドを実装しておけば、あらゆる入出力機能に対して、自由に実装コードを提供することが出来る。
また、利用する側も抽象化されているので、実装を意識することなくそれぞれのメソッドを利用することが出来る。
