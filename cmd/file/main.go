package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fileName := "/home/ivan/projects/testgo18/.debug/tls/cert.pem"
	if !checkFile(fileName) {
		fmt.Println("file no found")
		return
	}
	fmt.Println("file ok!")
}

func checkFile(fileName string) bool {
	_, errStat := os.Stat(fileName)
	if errors.Is(errStat, os.ErrNotExist) {
		return false
	}
	return true
}

func saveFile(fileName, data string) error {
	df, errCreateFile := os.Create(fileName)
	if errCreateFile != nil {
		return errCreateFile
	}
	defer df.Close()

	_, errWrite := df.WriteString(data)
	if errWrite != nil {
		return errWrite
	}
	return nil
}
