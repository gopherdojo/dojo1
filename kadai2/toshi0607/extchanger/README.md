# 第1回

## コマンドの実行

```
$ cd $GOPATH/src/github.com/toshi0607/gopher-dojo/extchanger/
$ go run main.go extension(from) extension(to) sourcedir

or

$ extchanger extension(from) extension(to) sourcedir
```

## GoDocの確認

```
$ godoc -http=:6060
$ curl http://localhost:6060/pkg/github.com/toshi0607/gopher-dojo/extchanger/converter/
```

# 第2回

## テストカバレッジ

```
go test -coverprofile=profile github.com/toshi0607/gopher-dojo/extchanger/converter
ok      github.com/toshi0607/gopher-dojo/extchanger/converter   1.390s  coverage: 90.6% of statements
```

## interface実装の一覧を出力（おまけ）

```
$ godoc -http ":4040" -analysys type
$ curl http://localhost:4040/pkg/io/#Writer

https://golang.org/docと異なり、implementsの項目が出力される
```