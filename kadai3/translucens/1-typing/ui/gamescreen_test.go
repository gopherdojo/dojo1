package ui

import (
	"strings"
	"testing"
)

func TestSingleTurn(t *testing.T) {
	type args struct {
		instr string
		word  string
	}
	tests := []struct {
		name       string
		args       args
		wantStrlen int
		wantScore  int
	}{
		{"perfect", args{"teststr", "teststr"}, 7, 7},
		{"miss1", args{"testst", "teststr"}, 7, 6},
		{"miss1", args{"test$tr", "teststr"}, 7, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan string)

			returnch := make(chan int)
			go func() {
				strlen, score := SingleTurn(ch, tt.args.word)
				returnch <- strlen
				returnch <- score
			}()
			ch <- tt.args.instr
			gotStrlen := <-returnch
			gotScore := <-returnch

			if gotStrlen != tt.wantStrlen {
				t.Errorf("SingleTurn() gotStrlen = %v, want %v", gotStrlen, tt.wantStrlen)
			}
			if gotScore != tt.wantScore {
				t.Errorf("SingleTurn() gotScore = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}

func TestTimeup(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name       string
		args       args
		wantStrlen int
		wantScore  int
	}{
		{"timeout", args{"teststr"}, 7, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan string)

			returnch := make(chan int)
			go func() {
				strlen, score := SingleTurn(ch, func() string { return tt.args.word })
				returnch <- strlen
				returnch <- score
			}()
			gotStrlen := <-returnch
			gotScore := <-returnch

			if gotStrlen != tt.wantStrlen {
				t.Errorf("SingleTurn() gotStrlen = %v, want %v", gotStrlen, tt.wantStrlen)
			}
			if gotScore != tt.wantScore {
				t.Errorf("SingleTurn() gotScore = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}

func TestStrInput(t *testing.T) {

	reader := strings.NewReader("teststring\n")
	ch := strinput(reader)

	gotstring := <-ch

	if gotstring != "teststring" {
		t.Errorf("strinput(r io.Reader) gotstring = %v", gotstring)
	}

}
