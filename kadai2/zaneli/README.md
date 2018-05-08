## 課題2-2

画像変換のテスト

### 使い方

GitHub から課題1のソースコードをダウンロードする。

```sh
> ./dl_src.sh
./dl_src.sh
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   966  100   966    0     0   3033      0 --:--:-- --:--:-- --:--:--  3037
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   897  100   897    0     0   2770      0 --:--:-- --:--:-- --:--:--  2768
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1059  100  1059    0     0   3337      0 --:--:-- --:--:-- --:--:--  3340
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1078  100  1078    0     0   3652      0 --:--:-- --:--:-- --:--:--  3654
```

テストを実行する。

```sh
> cd image-converter/converter/

> go test
PASS
ok  	_/Users/zaneli/ws/dojo1/kadai2/zaneli/image-converter/converter	1.461s

> go test --cover
PASS
coverage: 90.5% of statements
ok  	_/Users/zaneli/ws/dojo1/kadai2/zaneli/image-converter/converter	1.460s
```
