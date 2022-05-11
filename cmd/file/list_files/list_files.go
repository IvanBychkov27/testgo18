// http://прохоренок.рф/pdf/go/ch15-go-perebor-obyektov-raspolozhennykh-kataloge.html?
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Node struct {
	Files []string
}

func NewNode() *Node {
	return &Node{[]string{}}
}

func main() {
	//path, err := os.UserHomeDir()
	//if err != nil {
	//	fmt.Println("error user home dir:", err.Error())
	//	return
	//}
	//fmt.Println(path)

	node := NewNode()
	listDir(node, "/home/ivan/Iv/mp3", ".mp3")

	fmt.Println("files:", len(node.Files))
}

func listDir(node *Node, path, fileExtension string) {
	err := filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		if wPath == path {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if fileExtension != "" {
			if strings.Contains(wPath, fileExtension) {
				node.Files = append(node.Files, wPath)
			}
			return nil
		}

		node.Files = append(node.Files, wPath)
		return nil
	})

	if err != nil {
		fmt.Println("error walk:", err.Error())
	}
}

// =========================================
func listDirByWalk(path string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		// Обход директории без вывода
		if wPath == path {
			return nil
		}

		// Если данный путь является директорией, то останавливаем рекурсивный обход
		// и возвращаем название папки
		if info.IsDir() {
			fmt.Printf("[%s]\n", wPath)
			return filepath.SkipDir
		}

		// Выводится название файла
		if wPath != path {
			fmt.Println(wPath)
		}
		return nil
	})
}

func listDirByReadDir(path string) {
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("error read dir:", err.Error())
		return
	}
	for _, val := range lst {
		if val.IsDir() {
			fmt.Printf("[%s]\n", val.Name())
		} else {
			fmt.Println(val.Name())
		}
	}
}
