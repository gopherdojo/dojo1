## goTypingGame
タイピングゲームのGo実装

## Features
- [x] 標準出力に英単語を出す（出すものは自由）
  - [x] 標準出力に定義した英単語を順番に表示する
- [x] 標準入力から1行受け取る
  - [x] 入力を受け取った場合、表示に使用した単語と一致するか比較し、結果を出力する
- [x] 制限時間内に何問解けたか表示する
  - [x] ３０秒のタイマーをセットし、タイムアップした場合強制的にゲームの結果を表示して終了

## How to use

```
$ go get github.com/yusukemisa/goTypingGame

$ go install github.com/yusukemisa/goTypingGame

$ goTypingGame
```