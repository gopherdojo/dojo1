package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/aakira/jpg2png/image"
	"github.com/aakira/jpg2png/file"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Illeagal arguments")
		return
	}

	imagePaths, err := file.FindFiles(os.Args[1], ".jpg")
	if err != nil {
		fmt.Printf("File I/O error: %v\n", err)
		return
	}

	for _, path := range imagePaths {
		dstPath := strings.TrimSuffix(path, filepath.Ext(path)) + ".png"

		src, srcError := image.ToImageFile(path)
		dst, dstError := image.ToImageFile(dstPath)

		if srcError != nil || dstError != nil {
			fmt.Printf("Src error: %v, Dst error: %v", srcError, dstError)
		}
		image.ConvertImage(src, dst)
	}
}
