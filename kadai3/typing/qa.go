package typing

type QA struct {
	Good      int
	Bad       int
	Counter   int
	Questions []string
}

// NewQA -
func NewQA(questions []string) *QA {
	return &QA{
		Good:      0,
		Bad:       0,
		Counter:   -1,
		Questions: questions,
	}
}

// MakeQuestion makes question from given English word list.
func (qa *QA) MakeQuestion() string {
	// 出題範囲数を超えたら初めに戻る
	if len(qa.Questions) <= qa.Counter+1 {
		qa.Counter = 0
	} else {
		qa.Counter++
	}

	return qa.Questions[qa.Counter]
}

// CheckAnswer returns whether or not it matches the answer of the input.
func (qa *QA) CheckAnswer(question, answer string) bool {
	if question == answer {
		qa.Good++
		return true
	} else {
		qa.Bad++
		return false
	}
}
