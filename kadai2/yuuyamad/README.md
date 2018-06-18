# 課題2

## 課題2-1
# io.Readerとio.Writerについて調べてみよう  
### 標準パッケージでどのように使われているか  
io.Readerとio.Writerがinterfaceとして下記のように定義されている  
- io.Writer  
https://github.com/golang/go/blob/master/src/io/io.go#L90-L92  
- io.Reader  
https://github.com/golang/go/blob/master/src/io/io.go#L77-L79  

fileアクセスやインターネットアクセス等の処理を抽象化するために使われている  
os/file.goでは下記のようにio.Readerとio.Writerのインターフェースをみたすように実装されている  
 - Read  
https://github.com/golang/go/blob/master/src/os/file.go#L104-L110  

 - Write  
https://github.com/golang/go/blob/master/src/os/file.go#L141-L160  

### io.Readerとio.Writerがあることで、どういう利点があるのか具体例を挙げて考えてみる  
ReadやWriteといったinterfaceに定義されているメソッドを呼ぶ時に同じ処理をしたい時に型によって呼ぶメソッドを気にしなくて良い。  

io.Writerで定義されいるWriteメソッドはバイト列 p を書き込み、書き込んだバイト数 n と、エラーが起きた場合はそのエラーerrorを返すようになっている  
```
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

io.Writerを満たしていればfile出力でも標準出力でもバッファへの書き込みでも下記のように同様の呼び出しかたで呼び出せることが利点です  
```
file.Write([]byte("hogehoge"))
os.Stdout.Write([]byte("hogehoge"))
buffer.Write([]byte("hogehoge"))
```
