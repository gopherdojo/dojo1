## 第二回宿題 課題1
io.Reader,io.Writerについて調べよう

* 標準パッケージでどのように使われているか
* io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

#### 標準入出力(os.Stdin,os.Stdout,あとstderrも)ではどのように使われているか
標準入出力は対応するfd番号をもつ"/dev/std*"という名前のFile型構造体として扱われている。
```file.go
var (
    Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin") 
    Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```
#### ioパッケージ(Copyなど)ではどのように使われているか
Copy(io.Writer, io.Reader)という形になっている。
```
	//標準入力 > 標準出力
	if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}
```
上記のように標準入力と標準出力を指定するだけではなく
インターフェースを満たすものであれば以下のように何でも同じように扱うことができる。
```
	//メモリ > 標準出力
	if _, err := io.Copy(os.Stdout, bytes.NewReader([]byte("ほげええええええええ"))); err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}

	//ネットワーク > 標準出力
	resp, err := http.Get("https://yusukemisa.github.io/my-app5/")
	if err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}
```
上記のようにコマンドラインで発生する入出力だけでなくネットワーク越しの入出力も
ReaderとWriterという形で抽象化されていることにより同じ操作で同じことができる。

また、インターフェースを満たす者同士であれば入出力の組み合わせを自由に変えられる
ようになっておりとても柔軟性が高い。

