package jinja_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gopherdojo/dojo1/kadai4/execjosh/internal/jinja"
)

func TestGetBlessing(t *testing.T) {
	expectedValues := []struct {
		now      time.Time
		seed     int64
		expected []string
	}{
		{time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC), 1, []string{"大吉", "大吉", "大吉"}},
		{time.Date(2018, time.January, 2, 0, 0, 0, 0, time.UTC), 1, []string{"大吉", "大吉", "大吉"}},
		{time.Date(2018, time.January, 3, 0, 0, 0, 0, time.UTC), 1, []string{"大吉", "大吉", "大吉"}},

		{time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC), 2, []string{"大吉", "大吉", "大吉"}},
		{time.Date(2018, time.January, 2, 0, 0, 0, 0, time.UTC), 2, []string{"大吉", "大吉", "大吉"}},
		{time.Date(2018, time.January, 3, 0, 0, 0, 0, time.UTC), 2, []string{"大吉", "大吉", "大吉"}},

		{time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC), time.Now().UnixNano(), []string{"大吉", "大吉", "大吉"}},
		{time.Date(2018, time.January, 2, 0, 0, 0, 0, time.UTC), time.Now().UnixNano(), []string{"大吉", "大吉", "大吉"}},
		{time.Date(2018, time.January, 3, 0, 0, 0, 0, time.UTC), time.Now().UnixNano(), []string{"大吉", "大吉", "大吉"}},

		{time.Date(2018, time.May, 4, 0, 0, 0, 0, time.UTC), 1234567890, []string{"末小吉", "末吉", "末小吉", "末凶"}},
		{time.Date(2018, time.May, 17, 0, 0, 0, 0, time.UTC), 123567890, []string{"平", "平", "半吉", "末吉"}},
	}

	for _, tc := range expectedValues {
		tc := tc
		j := jinja.New(tc.seed)
		for _, expected := range tc.expected {
			expected := expected
			j := j
			actual := j.GetBlessing(tc.now)
			if expected != actual {
				fmt.Printf("expected '%v' but got '%v'\n", expected, actual)
				t.Fail()
			}
		}
	}
}
