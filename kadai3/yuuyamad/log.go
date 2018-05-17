package main

import (
	"log"
	"os"
)

func logError(err error){
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
}