package game

//出題する単語集
var words = []string{
	"golang",
	"gopher",
	"car",
	"cat",
	"dog",
}

//GetQuestion provide word for The Game
//TODO:ひとまず配列を順番に出すだけ
//出題ロジックを拡張できるように関数として作っておく
func getQuestion() []string {
	return words
}
