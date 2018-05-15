# go concurrent downloader

## Command Line Options

### -p <num>
split ratio to download file

### -o <filename>
output file to <filename>

## example
./yuuyamad -p 20 -o hoge.tar.xz "https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.6.3.tar.xz"


## 疑問点
requestsメソッドをgoroutineで実行する時にcontextをわたしているのだが、１つのgoroutineがtimeoutしたりキャンセル処理された時に
同じcontextを渡して実行されているgoroutineにもキャンセルが伝播してcontextのDoneチャネルが返るのかなと考えて実装しました

timeoutが他のgoroutineに伝播しているかというのは上手く検証できず、、自信がないのでもし認識間違いなどあれば教えて頂きたいです。

