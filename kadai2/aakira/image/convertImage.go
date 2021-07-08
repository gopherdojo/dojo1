package image

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Convert to destination image from source file.
func ConvertImage(srcImage Image, dstImage Image) {

	if srcImage == nil || dstImage == nil {
		fmt.Println("File not found.")
		return
	}

	srcPath := srcImage.GetPath()

	sf, err := os.Open(srcPath)
	if err != nil {
		fmt.Printf("File not found : %s \n", srcPath)
		return
	}
	defer sf.Close()

	dstPath := strings.TrimSuffix(srcPath, filepath.Ext(srcPath)) + dstImage.GetExt()

	df, err := os.Create(dstPath)
	if err != nil {
		fmt.Printf("File not create : %s \n", dstPath)
		return
	}
	defer df.Close()

	img, err := srcImage.Decode(sf)
	if err != nil {
		fmt.Printf("Decode error : %s \n", err)
		return
	}
	dstImage.Encode(df, img)

	if err != nil {
		fmt.Printf("Convert error: %s \n", err)
		return
	}
	fmt.Printf("Convert from [%s] to [%s] \n", srcPath, dstPath)
}
