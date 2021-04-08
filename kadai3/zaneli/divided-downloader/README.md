## 課題3-2

分割ダウンローダー

### 使い方

引数にダウンロード対象のURLを指定する。(必須)
並列度を `-p` オプションで、ダウンロード先ディレクトリを `-d` オプションで指定する。

```go
go run main.go https://dl.google.com/go/go1.10.2.src.tar.gz
```

```go
go run main.go -p=3 https://dl.google.com/go/go1.10.2.src.tar.gz
```

```go
go run main.go -d=/tmp/ https://dl.google.com/go/go1.10.2.src.tar.gz
```
