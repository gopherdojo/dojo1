package main

import (
	"testing"
	"reflect"
)

func TestReadfile(t *testing.T){

	var tests = []struct {
		filename string
		word     []string
	}{
		{"testdata/hoge.txt", []string{"hoge", "moga", "hoge", "moga", "moga", "fuga"}},
		{"testdata/fuga.txt", []string{"hogehoge", "mogamoga", "fugafuga", "hogehoge"}},
		{"testdata/moga.txt", []string{"hoga", "hoge", "fuga"}},
		{"testdata/unknown", []string{}},

	}
	for _, test := range tests {
		ret, err := readfile(test.filename)
		if test.filename == "testdata/unknown" && err == nil {
			t.Error(`readfile(testdata/unknown)`)
		}

		if test.filename != "testdata/unknown" && !reflect.DeepEqual(test.word, ret){
			t.Error(`readfile(%q)`, test.filename)
		}

	}

}

func Testquiz(t *testing.T){

}
