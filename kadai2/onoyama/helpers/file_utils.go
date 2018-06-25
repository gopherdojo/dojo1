package helpers

import (
	//"fmt"
  "os"
	"io/ioutil"
	"path/filepath"
  "strings"
)

// 指定ディレクトリをクロールして、ファイル情報をFileSpec型に格納
func DirWalker(dir string) []FileSpec {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
		//return ("ディレクトリ参照に失敗しました。", err)
	}

	var specs []FileSpec
	for _, file := range files {
		if file.IsDir() {
			specs = append(specs, DirWalker(filepath.Join(dir, file.Name()))...)
			continue
		}
    specs = append(specs, FileSpec{
      DirPath: dir,
      FileName: file.Name(),
      FileExt: filepath.Ext(file.Name()),
      BaseName: strings.Replace(file.Name(), filepath.Ext(file.Name()), "", 1),
    })
	}
	return specs
}

// FilesSpecから書き出し画像のパス等を生成して、ConvertSpec型に格納
func MakeConvertSpec(files []FileSpec, destPath string, toFormat string) []ConvertSpec {
  var specs []ConvertSpec
  for _, f := range files {
    ext := checkExt(f.FileExt)
    if ext {
      targetExt := TargetExt["."+toFormat]
      specs = append(specs, ConvertSpec{
        Src: f.DirPath + "/" + f.FileName,
        //Dst: destPath + "/" + f.DirPath + "/" + f.BaseName + targetExt,
        Dst: destPath + "/" + f.BaseName + targetExt,
        Format: toFormat,
      }  )
    } else {
      specs = append(specs, ConvertSpec{
        Src: f.DirPath + "/" + f.FileName,
        Dst: "",
        Format: "",
      })
    }
  }
  return specs
}

// ファイル拡張子が対応可能な画像のものかをチェック
func checkExt(extStr string) bool {
  result := false
  for _, ext := range PermitExt {
    if ext == strings.ToLower(extStr) {
      result = true
    }
  }
  return result
}

// ディレクトリやファイルの存在確認
func dirCheck(src string) bool {
  _, err := os.Stat(src)
  if err != nil {
    if os.IsNotExist(err) {
      return false
    }
  }
  return true
}

// ディレクトリの存在確認をして、なければディレクトリを作成
func dirCheckAndMkdir(src string) bool {
  srcDir := filepath.Dir(src)
  dirExt := dirCheck(srcDir)
  if !dirExt {
    err := os.MkdirAll(srcDir, os.ModePerm)
    if err != nil {
      panic(err)
    }
  }
  return true
}

