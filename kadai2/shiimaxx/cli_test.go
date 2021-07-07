package main

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var cases = []struct {
	name         string
	args         string
	outputs      []string
	outputFormat string
}{
	{
		name: "jepg to png",
		args: "image-convert testdata/jpeg",
		outputs: []string{
			"testdata/jpeg/icon-001.png",
			"testdata/jpeg/dir-001/icon-002.png",
			"testdata/jpeg/dir-002/icon-003.png",
			"testdata/jpeg/dir-002/dir-002-001/icon-004.png",
		},
		outputFormat: "png",
	},
	{
		name: "jepg to gif",
		args: "image-convert -s jpg -d gif testdata/jpeg",
		outputs: []string{
			"testdata/jpeg/icon-001.gif",
			"testdata/jpeg/dir-001/icon-002.gif",
			"testdata/jpeg/dir-002/icon-003.gif",
			"testdata/jpeg/dir-002/dir-002-001/icon-004.gif",
		},
		outputFormat: "gif",
	},
	{
		name: "png to jpeg",
		args: "image-convert -s png -d jpeg testdata/png",
		outputs: []string{
			"testdata/png/icon-001.jpg",
			"testdata/png/dir-001/icon-002.jpg",
			"testdata/png/dir-002/icon-003.jpg",
			"testdata/png/dir-002/dir-002-001/icon-004.jpg",
		},
		outputFormat: "jpeg",
	},
	{
		name: "png to gif",
		args: "image-convert -s png -d gif testdata/png",
		outputs: []string{
			"testdata/png/icon-001.gif",
			"testdata/png/dir-001/icon-002.gif",
			"testdata/png/dir-002/icon-003.gif",
			"testdata/png/dir-002/dir-002-001/icon-004.gif",
		},
		outputFormat: "gif",
	},
	{
		name: "gif to jpeg",
		args: "image-convert -s gif -d jpeg testdata/gif",
		outputs: []string{
			"testdata/gif/icon-001.jpg",
			"testdata/gif/dir-001/icon-002.jpg",
			"testdata/gif/dir-002/icon-003.jpg",
			"testdata/gif/dir-002/dir-002-001/icon-004.jpg",
		},
		outputFormat: "jpeg",
	},
	{
		name: "gif to png",
		args: "image-convert -s png -d gif testdata/gif",
		outputs: []string{
			"testdata/gif/icon-001.gif",
			"testdata/gif/dir-001/icon-002.gif",
			"testdata/png/dir-002/icon-003.gif",
			"testdata/png/dir-002/dir-002-001/icon-004.gif",
		},
		outputFormat: "gif",
	},
}

func TestRun_ImageConvert(t *testing.T) {
	t.Helper()
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			args := strings.Split(c.args, " ")
			status := cli.Run(args)
			if status != ExitCodeOK {
				t.Errorf("expected %d to eq %d", status, ExitCodeOK)
			}
			for _, o := range c.outputs {
				_, err := os.Stat(o)
				if os.IsNotExist(err) {
					t.Error(err)
				}
				f, err := os.Open(o)
				if err != nil {
					t.Fatal(err)
				}
				defer f.Close()
				_, format, err := image.DecodeConfig(f)
				if err != nil {
					t.Fatal(err)
				}
				if format != c.outputFormat {
					t.Errorf("expected %s to eq %s", format, c.outputFormat)
				}

			}
		})
	}
}

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("image-convert -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprintf("image-convert version %s", Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun_noArguments(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := []string{"image-convert"}

	status := cli.Run(args)
	if status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}

	expected := "Missing arguments\n"
	if !strings.EqualFold(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun_fileNotExists(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("image-convert dummy_file", " ")

	status := cli.Run(args)
	if status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}

	expected := "dummy_file: No such file or directory\n"
	if !strings.EqualFold(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun_isNotDir(t *testing.T) {
	t.Helper()
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	tempfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempfile.Name())

	args := strings.Split(fmt.Sprintf("image-convert %s", tempfile.Name()), " ")

	status := cli.Run(args)
	if status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}

	expected := fmt.Sprintf("%s: Is a not directory\n", tempfile.Name())
	if !strings.EqualFold(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}
