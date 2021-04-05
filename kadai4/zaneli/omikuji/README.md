## 課題4

おみくじAPI

### 使い方

ポートを指定しないで起動する。(デフォルト8080ポートが使用される。)

```sh
go run main.go
```

ポートを指定して起動する。

```sh
go run main.go 8081
```

アクセスする。

```
> curl --dump-header - "http://localhost:8080/"
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sat, 02 Jun 2018 17:52:03 GMT
Content-Length: 52

{"result":"吉","date":"2018-06-03T02:52:03+09:00"}
```

テストを実行する。

```
> go test ./omikuji
ok  	_/Users/zaneli/ws/dojo1/kadai4/zaneli/omikuji/omikuji	0.026s
```

```
> go test ./omikuji --cover
ok  	_/Users/zaneli/ws/dojo1/kadai4/zaneli/omikuji/omikuji	0.026s	coverage: 76.2% of statements
```
