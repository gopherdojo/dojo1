package cfile

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// ConvDirInfo 変換ディレクトリ情報
type ConvDirInfo struct {
	fs        afero.Fs
	inputDir  string
	outputDir string
}

// ConvFile 変換ファイルパス情報
type ConvFile struct {
	InputFile  string
	OutputFile string
}

// NewConvDirInfo 新規変換ディレクトリ情報作成
func NewConvDirInfo(fs afero.Fs, inputPath string) (ConvDirInfo, error) {
	ci := ConvDirInfo{}
	ci.fs = fs

	// inputPathチェック
	info, err := ci.fs.Stat(inputPath)
	if err != nil {
		return ci, err
	}
	if !info.IsDir() {
		return ci, errors.New("not a directory (input)")
	}
	ci.inputDir = inputPath

	return ci, nil
}

// SetOutputDir 出力ディレクトリ指定
func (ci *ConvDirInfo) SetOutputDir(outputPath string) error {
	// outputPathチェック
	if outputPath == "" {
		return nil
	}
	info, err := ci.fs.Stat(outputPath)
	if err != nil {
		if err := ci.fs.MkdirAll(outputPath, 0755); err != nil {
			return err
		}
	} else {
		if !info.IsDir() {
			return errors.New("not a directory (output)")
		}
	}
	ci.outputDir = outputPath
	return nil
}

// GetFiles 変換ファイルパスセット取得
func (ci *ConvDirInfo) GetFiles(inputExt string, outputExt string) ([]ConvFile, error) {
	var res []ConvFile
	err := afero.Walk(ci.fs, ci.inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			currentExt := strings.TrimLeft(filepath.Ext(info.Name()), ".")
			if currentExt == inputExt {
				op, err := ci.getOutputPath(path)
				if err != nil {
					return err
				}
				ope := filepath.Ext(op)
				op = fmt.Sprintf("%s.%s", op[0:len(op)-len(ope)], outputExt)
				cf := ConvFile{
					InputFile:  path,
					OutputFile: op,
				}
				res = append(res, cf)
			}
		}
		return nil
	})
	return res, err
}

// getOutputPath 出力パス取得
func (ci *ConvDirInfo) getOutputPath(inputPath string) (string, error) {
	outputDir := ci.outputDir
	if outputDir == "" {
		outputDir = ci.inputDir
	}
	baseInput, err := filepath.Abs(ci.inputDir)
	if err != nil {
		return "", err
	}
	baseOutput, err := filepath.Abs(outputDir)
	if err != nil {
		return "", err
	}
	inputFullPath, err := filepath.Abs(inputPath)
	if err != nil {
		return "", err
	}
	outputFullPath := fmt.Sprintf(
		"%s%s",
		baseOutput,
		inputFullPath[len(baseInput):len(inputFullPath)],
	)
	return outputFullPath, nil
}
