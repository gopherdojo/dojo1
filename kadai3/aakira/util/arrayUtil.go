package util

import (
	"math/rand"
	"time"
)

func ShuffleList(src []string) []string {
	rand.Seed(time.Now().UnixNano())

	dest := make([]string, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}
