/* This command convert jpg image to png, gif, tif, etc.

Command options

- Directory name

You must place jpg images in this directory. All jpg image in this direcotory is to be converted.

- Extenstion

A extension name you want images to converted.

Example

./main images png

*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"supermarine1377/processes"
)

func main() {
	var dirName string
	var extention string
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("no argument passed, exiting...")
		os.Exit(1)
	}
	dirName = args[1]
	if len(args) >= 3 {
		extention = args[2]
	} else {
		extention = "png"
	}
	if _, err := ioutil.ReadDir(dirName); err != nil {
		log.Printf("no dir %s found, exiting...", dirName)
		os.Exit(1)
	}
	images, err := processes.GetImages(dirName)
	if err != nil {
		log.Println("error occured during reading images", err)
	}
	for _, image := range images {
		if err := processes.Convert(image, extention); err != nil {
			log.Println(err)
		}
		log.Printf("finished converting %s", image.FileName)
	}
}
