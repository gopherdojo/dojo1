package typing

import (
	"fmt"
)

var testWords = []string{
	"abcdefg",
	"hijk",
	"lmnopqastu",
}

var testAnswers = []string{
	"abcdefg",
	"hijl", // typo
	"lmnopqastu",
}

func ExampleQA_MakeQuestion() {
	qa := NewQA(testWords)

	for i := 0; i < len(testWords)+2; i++ {
		fmt.Println(qa.MakeQuestion())
	}

	// Output:
	// abcdefg
	// hijk
	// lmnopqastu
	// abcdefg
	// hijk
}

func ExampleQA_CheckAnswer() {
	qa := NewQA(testWords)
	for _, a := range testAnswers {
		q := qa.MakeQuestion()
		fmt.Println(qa.CheckAnswer(q, a))
	}

	// Output:
	// true
	// false
	// true
}
