package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"reflect"
	"time"
	"context"

	"github.com/dmgk/faker"
	"github.com/fatih/color"
)

var (
	fakerMethods = []string{"Adjective", "Noun", "Verb", "IngVerb"}
	goodMessage  = "✓ GOOD!"
	badMessage   = "✗ BAD!"
	timeout      = 60 * time.Second
	wordCount    = 0
	goodCount    = 0
	badCount     = 0
	streaks      = 0
)

func main() {
	fmt.Print("press enter to start a minute typing! are you ready?")
	bufio.NewScanner(os.Stdin).Scan()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	typing(ctx)
	os.Exit(0)
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func typing(ctx context.Context) {
	ch := input(os.Stdin)
	for {
		w := word()
		fmt.Println("> " + w)
		select {
		case v, ok := <-ch:
			if ok {
				output(w, v)
			} else {
				return
			}
		case <-ctx.Done():
			fmt.Printf("\n :\n :\ntimeup! %v words correct out of %v.\n", goodCount, wordCount)
			return
		}
	}
}

func word() string {
	fh := faker.Hacker()
	methodName := fakerMethods[rand.Intn(len(fakerMethods))]
	rvm := reflect.ValueOf(fh).MethodByName(methodName)
	rv := rvm.Call([]reflect.Value{})
	if w, ok := rv[0].Interface().(string); ok {
		return w
	}
	// このエラーは main まで持っていくべきだろうか。ここで os.Exit するのは違和感がある
	fmt.Printf("word generation error. fakerMethod: %v\n", methodName)
	os.Exit(1)
	return ""
}

func output(w, v string) {
	wordCount += 1
	if w == v {
		goodCount += 1
		streaks += 1
		if streaks > 1 {
			color.Green("%v %v streaks!", goodMessage, streaks)
		} else {
			color.Green("%v", goodMessage)
		}
	} else {
		badCount += 1
		streaks = 0
		color.Red("%v", badMessage)
	}
	fmt.Println("")
}
