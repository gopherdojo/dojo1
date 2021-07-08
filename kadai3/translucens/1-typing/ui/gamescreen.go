package ui

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/translucens/dojo1/kadai3/translucens/1-typing/typinggame"
)

const (
	turncount    = 5
	warmuptime   = 3
	timeperchar  = 1
	timetrialsec = 30
)

var (
	yellow       = color.New(color.FgYellow).SprintFunc()
	red          = color.New(color.FgRed).SprintFunc()
	blue         = color.New(color.FgHiBlue).SprintFunc()
	whiteBgcyan  = color.New(color.FgHiWhite).Add(color.BgCyan).Add(color.Bold).SprintFunc()
	whiteBggreen = color.New(color.FgHiWhite).Add(color.BgGreen).SprintfFunc()
)

// high score and name

// MainScreen is mode selector
func MainScreen() {

	fmt.Println("  ______                 ______")
	fmt.Println(" /_  __/_  ______  ___  / ____/___")
	fmt.Println("  / / / / / / __ \\/ _ \\/ / __/ __ \\")
	fmt.Println(" / / / /_/ / /_/ /  __/ /_/ / /_/ /")
	fmt.Println("/_/  \\__, / .___/\\___/\\____/\\____/")
	fmt.Println("    /____/_/")

	strchan := strinput(os.Stdin)
	defer close(strchan)

	for {
		fmt.Println("Select game mode: ")
		fmt.Printf("1: %d turns\n", turncount)
		fmt.Printf("2: Timetrial %d sec.\n", timetrialsec)
		fmt.Println("Other: Exit")
		fmt.Print(whiteBgcyan(">>> "))

		command, ok := <-strchan
		switch {
		case len(command) == 0 || !ok:
			return
		case command[0] == '1':
			printScore(TurnGame(strchan))
		case command[0] == '2':
			printScore(Timetrial(strchan))
		default:
			return
		}

	}

}

// TurnGame shows turn-ruled game screen for player
func TurnGame(strchan <-chan string) (int, int) {

	totalscore := 0
	totallength := 0

	for i := warmuptime; i > 0; i-- {
		fmt.Printf("%d...", i)
		time.Sleep(time.Second)
	}

	for i := 1; i <= turncount; i++ {
		fmt.Print("Ready...")
		time.Sleep(time.Second)

		fmt.Printf("Go!!\nTurn %d/%d\n", i, turncount)
		word := typinggame.RandomWord()
		lenstr, score := SingleTurn(strchan, word, time.Duration(len(word)*timeperchar)*time.Second)

		totallength += lenstr
		totalscore += score
	}
	return totallength, totalscore
}

// Timetrial shows time-based game screen
func Timetrial(strchan <-chan string) (int, int) {

	totallength, totalscore := 0, 0

	for i := warmuptime; i > 0; i-- {
		fmt.Printf("%d...", i)
		time.Sleep(time.Second)
	}
	fmt.Print("Ready...")
	time.Sleep(time.Second)
	fmt.Println("Go!!")

	endAt := time.Now().Add(timetrialsec * time.Second)

	for endAt.After(time.Now()) {

		word := typinggame.RandomWord()

		lenstr, score := SingleTurn(strchan, word, endAt.Sub(time.Now()))

		totallength += lenstr
		totalscore += score
	}

	return totallength, totalscore
}

func printScore(charcount, score int) {
	fmt.Println()
	fmt.Println(yellow(" ********************************** "))
	fmt.Printf("* Total Score: %d; Accuracy: %.1f%% *\n", score, float64(score)*100.0/float64(charcount))
	fmt.Println(yellow(" ********************************** "))
}

// SingleTurn shows typing object and returns word length and score
func SingleTurn(strchan <-chan string, correctstr string, timeout time.Duration) (int, int) {

	fmt.Println(whiteBggreen("### %s ### %d [sec.]", correctstr, timeout/time.Second))
	fmt.Print(">>> ")

	lenstr := len(correctstr)

	timer := time.NewTimer(timeout)

	for {
		select {
		case playerstr, ok := <-strchan:

			score := 0
			if ok {
				score = typinggame.CalcScore(correctstr, playerstr)

				if score == lenstr {
					fmt.Print(yellow("PERFECT! "))
				} else {
					fmt.Print(red("miss... "))
				}

				fmt.Printf("Got %d point !\n", score)
			}
			timer.Stop()

			return lenstr, score
		case _ = <-timer.C:
			fmt.Printf("\nTimeup !!\n")
			return lenstr, 0
		}
	}

}

func strinput(r io.Reader) chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		if err := s.Err(); err != nil {
			fmt.Println(err.Error())
		}
		// EOF
		close(ch)
	}()

	return ch
}
