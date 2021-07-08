package network

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	PathFound                = "found"
	PathWithoutContentLength = "withoutcontentlength"
	PathNotFound             = "notfound"
)

var testCondition int

var testhandler = http.HandlerFunc(
	func(writer http.ResponseWriter, req *http.Request) {

		switch testCondition {
		case 0:
			http.ServeFile(writer, req, "../testdata/1024")
		case 1:
			writer.WriteHeader(http.StatusOK)
		case 2:
			http.NotFound(writer, req)
		}
	})

func TestGetFileSize(t *testing.T) {
	tests := []struct {
		name      string
		condition int
		want      int64
		wantErr   bool
	}{
		{"200OK", 0, 1024, false},
		{"200OKwithoutContentLength", 1, -1, false},
		{"404NotFound", 2, 0, true},
	}

	testserver := httptest.NewServer(testhandler)
	defer testserver.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testCondition = tt.condition
			got, err := GetFileSize(testserver.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFileSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cutFileName(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"withdirectories", args{"/a/b/file.txt"}, "file.txt"},
		{"nodir", args{"/file.htm"}, "file.htm"},
		{"endsdir", args{"/a/dir/"}, "dir"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutFileName(tt.args.url); got != tt.want {
				t.Errorf("cutFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
