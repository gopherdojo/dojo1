package main

import (
	"fmt"
	"github.com/aakira/typinggame/file"
	"github.com/aakira/typinggame/game"
	"github.com/aakira/typinggame/util"
)

func main() {
	wordList := util.ShuffleList(file.ReadFile("wordList.txt", 100))
	point := game.Start(wordList, 60)
	fmt.Printf("Your score is %d.\n", point)
}
