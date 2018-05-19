package jinja

import (
	"math/rand"
	"sync"
	"time"
)

var blessings = []string{
	"大吉",
	"吉",
	"中吉",
	"小吉",
	"半吉",
	"末吉",
	"末小吉",
	"平",
	"凶",
	"小凶",
	"半凶",
	"末凶",
	"大凶",
}
var blessingCount = len(blessings)

// A Jinja that provides random omikuji
type Jinja struct {
	lk  sync.Mutex
	rng *rand.Rand
}

// New creates and initializes a new Jinja
func New(seed int64) *Jinja {
	return &Jinja{
		rng: rand.New(rand.NewSource(seed)),
	}
}

// GetBlessing returns an blessing based on the current time.
// For 1/1-1/3, it returns always "大吉".
// For all other dates, it returns a random blessing.
func (m *Jinja) GetBlessing(now time.Time) string {
	if now.YearDay() <= 3 {
		return "大吉"
	}

	m.lk.Lock()
	n := m.rng.Intn(blessingCount)
	m.lk.Unlock()

	return blessings[n]
}
