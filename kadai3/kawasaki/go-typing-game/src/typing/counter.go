package typing

// 引数として受け取った受信チャネルから正解のboolをうけとり
// その数をカウントする
// 戻り値のカウントはintのポインター
func counter(ch <-chan bool) *int {
	count := 0
	go func() {
		for {
			if <-ch {
				count++
			}
		}
	}()
	return &count
}
