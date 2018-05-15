package split

import (
	"."
	"testing"
)

func TestCreateArray(t *testing.T) {
	test1 := split.CreateArray(15, 3)
	exp := []int{5, 10, 15}
	for i, x := range exp {
		if test1[i] != x {
			t.Fatalf("Error by split.CreateArray. %v", test1)
		}
	}
}
