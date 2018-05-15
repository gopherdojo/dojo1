package typinggame

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var words = [...]string{"break", "default", "func", "interface", "select",
	"case", "defer", "go", "map", "struct",
	"chan", "else", "goto", "package", "switch",
	"const", "fallthrough", "if", "range", "type",
	"continue", "for", "import", "return", "var"}

// RandomWord returns a random chosen word
func RandomWord() string {

	return words[rand.Intn(len(words))]
}
