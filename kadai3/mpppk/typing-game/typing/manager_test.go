package typing

import (
	"testing"
)

func TestManager(t *testing.T) {
	cases := []struct {
		questions []string
	}{
		{
			questions: []string{"foo"},
		},
		{
			questions: []string{"foo", "bar"},
		},
	}

	for _, c := range cases {
		manager := NewManager()
		manager.AddQuestions(c.questions)
		cnt := 0
		for manager.SetNewQuestion() {
			if !manager.ValidateAnswer(c.questions[cnt]) {
				t.Fatalf("%dth question must be %q but actually %q (questions: %q)",
					cnt+1,
					c.questions[cnt],
					manager.GetCurrentQuestion(),
					c.questions)
			}
			cnt++
		}

		if cnt != len(c.questions) {
			t.Fatalf("question num is %d, but cnt is %d", len(c.questions), cnt)
		}
	}
}
