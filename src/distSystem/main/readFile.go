package main

import (
	"os"
	"log"
)

func main() {
	basePath := "C:\\git\\GoWork\\src\\distSystem\\main\\kjv12.txt"
	templateFolder, err := os.Open(basePath)
	if err != nil {
		log.Fatal(err)
	}
	defer templateFolder.Close()

	outfile, err := os.Create(basePath+"0")
	if err != nil {
		log.Fatal("Split4: ", err);
	}
	outfile.Close()
}
