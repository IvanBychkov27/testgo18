// http://прохоренок.рф/pdf/go/ch15-go-funktsii-dlya-raboty-s-katalogami.html

// UserHomeDir() — позволяет получить строковое представление домашнего каталога
// Getwd() — позволяет получить строковое представление текущего рабочего каталога

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var dirNames []string

func main() {
	allDirs()
}

type DirDevice struct {
	Path  string
	Name  string
	Count int
}

func allDirs() {
	path, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error user home dir:", err.Error())
		return
	}
	fmt.Println(path)

	listDir(path, ".mp3")
	fmt.Println("files:", len(dirNames))

	dirs := make(map[string]int, 0)
	for _, d := range dirNames {
		name := filepath.Dir(d)
		if strings.Contains(name, ".local") {
			continue
		}
		dirs[name]++

		inc := 0
		name = strings.TrimPrefix(name, path)
		for {
			inc++
			name = filepath.Dir(name)
			if name == "/" || name == "." || inc > 30 {
				break
			}
			dirs[name]++
		}
	}

	fmt.Println("dirs :", len(dirs))

	dirInfo := make([]DirDevice, 0, len(dirs))
	for pathDir, count := range dirs {
		d := DirDevice{
			Path:  pathDir,
			Name:  strings.TrimPrefix(pathDir, path),
			Count: count,
		}
		dirInfo = append(dirInfo, d)
	}

	sort.Slice(dirInfo, func(i, j int) bool {
		return dirInfo[i].Name < dirInfo[j].Name
	})

	for _, v := range dirInfo {
		fmt.Printf("%4d - %s\n", v.Count, v.Name)
	}

}

func listDir(path, fileExtension string) {
	err := filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		if wPath == path {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if fileExtension != "" {
			if strings.Contains(wPath, fileExtension) {
				dirNames = append(dirNames, wPath)
			}
			return nil
		}

		dirNames = append(dirNames, wPath)
		return nil
	})

	if err != nil {
		fmt.Println("error walk:", err.Error())
	}
}
