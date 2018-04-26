# 課題2

## 課題2-1

io.Readerとio.Writerについて調べてみよう。

* 標準パッケージでどのように使われているか
```
基本的なRead/Write系の実装メソッドの引数/戻り値として、このインターフェイスが利用されているようでした。
抽象型を利用することでそれぞれの実装メソッドが汎用的に利用できるようになる印象です。
```

* io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
```
読み込み/書き込みの実装をケースによって切り替えられ、通常はディスク情報を読み書きする実装にして、テストの際はメモリ上でのみ読み書きする実装にする事ができる。
```

## 課題2-2

課題1のテストを作ってみて下さい。

* テストのしやすさを考えてリファクタリングしてみる
* テストのカバレッジを取ってみる
* テーブル駆動テストを行う
* テストヘルパーを作ってみる

### 備考
- ファイルシステムの抽象化を自前でやろうとしたんですが、結構深いことになり時間がかかりそうだったので、サードパーティのパッケージを利用しました→"github.com/spf13/afero"

### テストのカバレッジ
```
> go test -coverprofile=cfile github.com/xlune/dojo1/kadai2/tenntenn/imgconv/cfile                       
ok  	github.com/xlune/dojo1/kadai2/tenntenn/imgconv/cfile	0.025s	coverage: 81.6% of statements

> go test -coverprofile=cimage github.com/xlune/dojo1/kadai2/tenntenn/imgconv/cimage
ok  	github.com/xlune/dojo1/kadai2/tenntenn/imgconv/cimage	0.027s	coverage: 82.1% of statements
```
(※ファイルの作成失敗やイメージ変換失敗処理のコード辺りがカバーしきれていない状況です)