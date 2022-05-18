package main

import (
	"flag"
	"fmt"
)

func main() {
	fileNameRead := flag.String("read", "", "reading file")
	fileNameSave := flag.String("out", "", "saved file")
	flag.Parse()

	fmt.Println("fileNameRead:", *fileNameRead)
	fmt.Println("fileNameSave:", *fileNameSave)
}
