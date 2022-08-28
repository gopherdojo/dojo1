package lib

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/gopherdojo/dojo1/kadai4/matsuki/lib"
)

// TestMain called first
func TestMain(t *testing.T) {

	beforeTest()

	// test here
	// TestIndex(t)

}

func beforeTest() {
	rand.Seed(time.Now().UnixNano())
	lib.Verbos = true
}

func TestDrawFortuneSlip(t *testing.T) {
	f := lib.NewFortune(time.Now())
	slip, _ := f.DrawFortuneSlip()
	switch slip {
	case "大吉", "中吉", "吉", "凶", "大凶":
		log.Print("draw valid slip")
	default:
		t.Errorf("invalid slip=%v", slip)
	}
}

func TestDrawWhenNewyear(t *testing.T) {
	tm, err := time.Parse("2006-01-02", "2013-01-03")
	log.Println("test date=", tm, err)

	f := lib.NewFortune(tm)
	slip, isNd := f.DrawFortuneSlip()
	switch slip {
	case "大吉":
		log.Print("draw valid slip")
	case "中吉", "吉", "凶", "大凶":
		t.Errorf("today is newYear. invalid slip=%v", slip)
	default:
		t.Errorf("invalid slip=%v", slip)
	}

	if isNd != true {
		t.Errorf("want isNewYearDay is true but %v", isNd)
	}
}
