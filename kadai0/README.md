# 選考課題：go-head

* headコマンドをGoで実装
```
# help表示
go-head -h

# ファイル名を指定して表示
go-head -f testfile.txt

# 標準入力から表示
cat testfile.txt | go-head -n 20
```

## setup & build

* buildのために、gopathをとおすため、このディレクトリをGOPATHに設定
```
./setup/setup.sh
```
* buildしてコマンドを利用できるようにする
```
./build/build.s
```
