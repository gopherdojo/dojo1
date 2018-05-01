package main

import (
	"fmt"
	"os"
	"time"
	"bufio"
	"strconv"
	"flag"
)

func main() {

	var (
		seconds = flag.Duration("t", 10*time.Second, "timer set")
	)
	flag.Parse()
	args := flag.Args()

	//問題用のファイルを読み込み
	word := readfile(args[0])

	fmt.Println("Start typing GAME!!!")

	ch := quiz(word)
	ch2 := time.After(*seconds)

	j := 0

	END:
	for {
		select {
		case <-ch2:
			break END
		case _, ok := <-ch:
			if ok {
				j++
			}else{
				break END
			}
		}
	}
	s := strconv.Itoa(j)
	fmt.Println("Correct Answer is " + s)

}

func readfile(file string) []string {

	var word []string
	var fp *os.File
	var err error

	fp, err = os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var sc = bufio.NewScanner(fp);
	sc.Split(bufio.ScanWords)
	for sc.Scan(){
		word = append(word, sc.Text())
	}
	return word
}

func quiz(quizes []string) <-chan struct{}{
	ch := make(chan struct{})

	go func(){

		for _, quiz := range quizes {
			fmt.Print(quiz + ":")
			s:= bufio.NewScanner(os.Stdin)
			s.Scan()
			scan := s.Text()
			if quiz == scan{
				ch <- struct{}{}
			}

		}
		close(ch)

	}()
	return ch
}

