# コマンドの実行

```
$ cd $GOPATH/src/github.com/toshi0607/gopher-dojo/extchanger/
$ go run main.go extension(from) extension(to) sourcedir

or

$ extchanger extension(from) extension(to) sourcedir
```

# GoDocの確認

```
$ godoc -http=:6060
$ curl http://localhost:6060/pkg/github.com/toshi0607/gopher-dojo/extchanger/converter/
```