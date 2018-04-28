# Gopher道場 課題2
## image_converter]
JPEG画像をPNG画像に変換。またはPNG画像をJPEG画像に変換

## コマンド
image_converter -d=出力先フォルダ -f=変換後の画像形式(jpegまたはpng) 変換元フォルダ

Example:
    image_converter -d=./jpeg -f=png ./output/png
    「./jpeg」ディレクトリ内のJPEG画像をPNG画像に変換して、「./output/png」ディレクトリに出力

## helpersライブラリ
PACKAGE DOCUMENTATION

package helpers
    import "./helpers"

    [image_converter]コマンドの実行支援用パッケージ

VARIABLES

var PermitExt = []string{".jpeg", ".jpg", ".png"}
    変換可能なファイル拡張子

var TargetExt = map[string]string{".jpeg": ".jpg", ".jpg": ".jpg", ".png": ".png"}
    画像拡張子の統一用

FUNCTIONS

func BulkConvert(specs []ConvertSpec)
    ConvertSpecに格納された情報で、複数の画像を変換

func CheckParams() (string, string, string)
    コマンドの引数の確認

func ConvertImageFile(src string, dst string, toFormat string) error
    指定のフォーマットで画像を変換する

func DirWalker(dir string) []FileSpec
    指定ディレクトリをクロールして、ファイル情報をFileSpec型に格納

func MakeConvertSpec(files []FileSpec, destPath string, toFormat string) []ConvertSpec
    FilesSpecから書き出し画像のパス等を生成して、ConvertSpec型に格納

func PrintMsgAndDie(comment string, message error)
    コメント、エラー、使用例を表示して、異常終了

func PrintUsage()
    コマンドの使用例を表示

TYPES

type ConvertSpec struct {
    Src    string
    Dst    string
    Format string
}
    画像変換時の情報格納

type FileSpec struct {
    DirPath  string
    FileName string
    BaseName string
    FileExt  string
}
    取得したファイル情報を格納

## テスト
go test ./helpers -v
