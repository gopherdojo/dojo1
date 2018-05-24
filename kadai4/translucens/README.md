# 課題4

おみくじAPIを作ってみよう。

* JSON形式でおみくじの結果を返す
  * `result`に結果が入るようにしました
* 正月（1/1-1/3）だけ大吉にする
  * 正月の前後で境界値試験を記述しました
* ハンドラのテストを書いてみる
  * `TestHTTPHandler`に記述しました

使い方

```sh
$ go run omikujiserver.go
# In another terminal
$ curl --silent localhost:8080 | jq .
{
  "result": "吉",
  "date": "2018-05-25 02:35:02.0177204 +0900 DST m=+106.863404001"
}
```
