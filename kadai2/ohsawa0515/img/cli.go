// refer https://deeeet.com/writing/2014/12/18/golang-cli-test/
package img

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// 終了コード
const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitError
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

type Opt struct {
	dir  string
	from string
	to   string
}

// Run -
func (cli *CLI) Run(args []string) int {
	var opt Opt
	flags := flag.NewFlagSet("goimgconverter", flag.ContinueOnError)
	flags.SetOutput(cli.ErrStream)
	flags.StringVar(&opt.dir, "d", "", "path of conversion destination")
	flags.StringVar(&opt.from, "f", "", "image extension before conversion")
	flags.StringVar(&opt.to, "t", "", "image extension after conversion")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if err := opt.validationOpt(); err != nil {
		log.Println(err)
		return ExitCodeParseFlagError
	}

	if err := cli.walk(opt.dir, opt.from, opt.to); err != nil {
		log.Println(err)
		return ExitError
	}

	return ExitCodeOK
}

func (opt *Opt) validationOpt() error {
	switch opt.from {
	case "jpg", "jpeg", "png":
	default:
		return errors.Errorf("%s: specify extension `jpg` or `png`.", opt.to)
	}

	switch opt.to {
	case "jpg", "jpeg", "png":
	default:
		return errors.Errorf("%s: specify extension `jpg` or `png`.", opt.to)
	}

	if !isExistDir(opt.dir) {
		return errors.Errorf("%s: No such file or directory", opt.dir)
	}

	return nil
}

func isExistDir(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func (cli *CLI) walk(root, from, to string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !isTargetFile(path, from) {
			return nil
		}
		r, err := os.Open(path)
		if err != nil {
			return err
		}
		defer r.Close()

		out := getOutputFile(path, to)
		w, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer w.Close()

		if err := Convert(r, w, filepath.Ext(path)); err != nil {
			return err
		}

		return nil
	})
}
