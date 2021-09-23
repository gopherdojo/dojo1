package main

import (
	"fmt"

	"github.com/yusukemisa/goTypingGame/game"
)

func main() {
	fmt.Println("★★★★★★★★ TYPING GAME START! ★★★★★★★★")
	gameResult := game.StartWithTimer(30)
	fmt.Printf("\n%v\n", gameResult.GameOverReason)
	fmt.Printf("Your correct answer is %v\n", gameResult.CorrectNum)
	fmt.Println("★★★★★★★★ TYPING GAME END! ★★★★★★★★")
}
