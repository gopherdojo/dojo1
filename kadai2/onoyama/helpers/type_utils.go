package helpers

// 変換可能なファイル拡張子
var PermitExt = []string{".jpeg", ".jpg", ".png"}
// 画像拡張子の統一用
var TargetExt = map[string]string{".jpeg":".jpg", ".jpg":".jpg", ".png":".png"}

// 画像変換時の情報格納
type ConvertSpec struct {
  Src string 
  Dst string
  Format string
}

// 取得したファイル情報を格納
type FileSpec struct {
  DirPath string
  FileName string
  BaseName string
  FileExt string
} 
