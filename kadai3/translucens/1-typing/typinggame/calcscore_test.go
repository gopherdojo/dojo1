package typinggame

import (
	"testing"
)

func TestEditLength(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"same", args{"test String", "test String"}, 0},
		{"missing one char", args{"test String", "test Strin"}, 1},
		{"add one char", args{"test String", "test Stringg"}, 1},
		{"empty", args{"test String", ""}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EditLength(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("EditLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcScore(t *testing.T) {
	type args struct {
		correct   string
		userinput string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"correct", args{"string", "string"}, 6},
		{"small miss", args{"string", "strin"}, 5},
		{"small miss2", args{"string", "stringg"}, 5},
		{"middle miss", args{"string", "stri"}, 4},
		{"middle miss2", args{"string", "string12"}, 4},
		{"big miss", args{"string", ""}, 0},
		{"big miss2", args{"string", "string1234567"}, 0},
		{"short string", args{"go", "go"}, 2},
		{"short miss", args{"go", "g"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcScore(tt.args.correct, tt.args.userinput); got != tt.want {
				t.Errorf("CalcScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
