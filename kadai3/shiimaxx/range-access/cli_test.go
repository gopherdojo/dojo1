package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	var cases = []struct {
		name         string
		args         string
		expected     string
		isNormalCase bool
	}{
		{
			name:         "version flag",
			args:         "gurl -version",
			expected:     fmt.Sprintf("gurl version %s\n", Version),
			isNormalCase: true,
		},
		{
			name:         "no arguments",
			args:         "gurl",
			expected:     "missing arguments\n",
			isNormalCase: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
			cli := &CLI{outStream: outStream, errStream: errStream}
			status := cli.Run(strings.Split(c.args, " "))

			if c.isNormalCase {
				if status != ExitCodeOK {
					t.Errorf("expected %d to eq %d", status, ExitCodeOK)
				}
				if !strings.EqualFold(outStream.String(), c.expected) {
					t.Errorf("expected %q to eq %q", outStream.String(), c.expected)
				}
			} else {
				if status != ExitCodeError {
					t.Errorf("expected %d to eq %d", status, ExitCodeOK)
				}
				if !strings.EqualFold(errStream.String(), c.expected) {
					t.Errorf("expected %q to eq %q", errStream.String(), c.expected)
				}
			}
		})
	}
}
