package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	allDir         []string
	allFiles       []string
	allNameFiles   []string
	nameFile       string
	events         chan int
	idx            int
	checkboxRamdom bool
	fraction       float32
	timePlay       string
	stopPlay       chan struct{}
	countPlay      int
	itemSelected   int32
	dirItems       []string
	dirInfo        []DirDevice
	currentPath    string
)

type DirDevice struct {
	Path  string
	Name  string
	Count int
}

const (
	unknown = iota
	play
	stop
	next
	dClick
	autoNext
)

func main() {
	rand.Seed(time.Now().UnixMilli())

	allDir = make([]string, 0, 1000)
	allFiles = make([]string, 0, 1000)
	dirItems = make([]string, 0, 100)
	dirItems = append(dirItems, "...")

	events = make(chan int)
	stopPlay = make(chan struct{})

	getListFiles()

	go allDirs()

	go processing()

	wnd := g.NewMasterWindow("Player Iv", 400, 500, 0).
		RegisterKeyboardShortcuts(
			g.WindowShortcut{
				Key:      g.KeyEscape,
				Modifier: g.ModNone,
				Callback: exitFunc},
		)
	wnd.Run(loop)
}

func loop() {
	g.SingleWindowWithMenuBar().Layout(
		g.MenuBar().Layout(
			g.Menu("File").Layout(
				g.Menu("Dir").Layout(
					g.Combo("", dirItems[itemSelected], dirItems, &itemSelected).OnChange(comboChanged),
				),
				g.MenuItem("Exit").OnClick(exitFunc),
			),
		),
		g.Row(
			g.Label("Play:"),
			g.Label(nameFile),
		),
		g.ProgressBar(fraction).Size(g.Auto, 0),
		g.Label(timePlay),
		g.Row(
			g.Button("Play").OnClick(OnPlay),
			g.Button("Stop").OnClick(OnPlayStop),
			g.Button("Next").OnClick(OnPlayNext),
			g.Checkbox("random", &checkboxRamdom),
		),
		g.ListBox("ListBox1", allNameFiles).OnDClick(listDClick),
	)
}

func comboChanged() {
	path := dirInfo[itemSelected].Path
	getListAllFiles(path, ".mp3")
	listNameFiles(path)
	currentPath = path
}

func exitFunc() {
	stopPlayMp3()
	saveListAllFiles(currentPath)
	time.Sleep(time.Millisecond * 200)
	os.Exit(0)
}

func allDirs() {
	path, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error user home dir:", err.Error())
		return
	}

	getListDir(path, ".mp3")

	dirs := make(map[string]int, 0)
	for _, d := range allDir {
		name := filepath.Dir(d)
		if strings.Contains(name, ".local") {
			continue
		}
		dirs[name]++

		inc := 0
		for {
			inc++
			name = filepath.Dir(name)
			if name == "" || name == "/" || name == "." || inc > 30 {
				break
			}
			dirs[name]++
		}
	}
	allDir = nil

	dirInfo = make([]DirDevice, 0, len(dirs))
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

	dirItems = nil
	for _, v := range dirInfo {
		if v.Name == "" {
			v.Name = "all"
		}
		dirItems = append(dirItems, fmt.Sprintf("%4d - %s", v.Count, v.Name))
	}
}

// получаем список всех файлов mp3
func getListAllFiles(path, fileExtension string) {
	allFiles = nil
	err := filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		if wPath == path {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if fileExtension != "" {
			if strings.HasSuffix(wPath, fileExtension) && !strings.Contains(wPath, ".local") {
				allFiles = append(allFiles, wPath)
			}
			return nil
		}

		allFiles = append(allFiles, wPath)
		return nil
	})

	if err != nil {
		fmt.Println("error walk:", err.Error())
	}
}

// получаем список всех папок с файлами mp3
func getListDir(path, fileExtension string) {
	err := filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		if wPath == path {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if fileExtension != "" {
			if strings.HasSuffix(wPath, fileExtension) && !strings.Contains(wPath, ".local") {
				allDir = append(allDir, wPath)
			}
			return nil
		}

		allDir = append(allDir, wPath)
		return nil
	})

	if err != nil {
		fmt.Println("error walk get list dir:", err.Error())
	}
}

// получаем список всех файлов mp3 без пути к нему
func listNameFiles(path string) {
	allNameFiles = make([]string, 0, len(allFiles))
	for i, nameFile := range allFiles {
		name := fmt.Sprintf("%d - %s", i, strings.TrimPrefix(nameFile, path+"/"))
		allNameFiles = append(allNameFiles, name)
	}
}

// сохраняем path в файл prayer_iv.lst
func saveListAllFiles(path string) {
	df, errCreateFile := os.Create("prayer_iv.lst")
	if errCreateFile != nil {
		fmt.Println("error create file:", errCreateFile.Error())
		return
	}
	defer df.Close()

	_, errWrite := df.WriteString(path)
	if errWrite != nil {
		fmt.Println("error write file:", errWrite.Error())
		return
	}
}

// получаем список файлов который был при предыдущем выходе из программы
func getListFiles() {
	file, errReadFile := ioutil.ReadFile("prayer_iv.lst")
	if errReadFile != nil {
		return
	}
	path := string(file)
	getListAllFiles(path, ".mp3")
	listNameFiles(path)
}

func processing() {
	for {
		select {
		case event := <-events:
			switch event {
			case play, dClick: // play, DClick
				stopPlayMp3()
			case stop: // stop
				stopPlayMp3()
				timePlay = ""
				fraction = 0
				continue
			case next: // next
				stopPlayMp3()
				if checkboxRamdom {
					idx = rand.Intn(len(allFiles))
				} else {
					idx++
				}
			case autoNext: // auto next
				if checkboxRamdom {
					idx = rand.Intn(len(allFiles))
				} else {
					idx++
				}
			}
		}

		time.Sleep(time.Millisecond * 200)

		if idx > len(allFiles)-1 {
			idx = 0
		}

		if len(allNameFiles) == 0 {
			continue
		}
		nameFile = allNameFiles[idx] // имя выбранного файла - выводим на экран

		go playMp3()
	}
}

func playMp3() {
	countPlay++
	defer func() { countPlay-- }()

	file, errOpen := os.Open(allFiles[idx])
	if errOpen != nil {
		log.Fatal(errOpen)
	}
	defer file.Close()

	streamer, format, errMP3Decode := mp3.Decode(file)
	if errMP3Decode != nil {
		log.Fatal(errMP3Decode)
	}
	defer streamer.Close()

	sr := format.SampleRate * 1 // скорость воспроизведения файла (*2 - ускоряет в 2 раза!)
	errInit := speaker.Init(sr, sr.N(time.Second/10))
	if errInit != nil {
		log.Fatal(errInit)
	}

	ctrl := &beep.Ctrl{Streamer: beep.Loop(1, streamer), Paused: false} // -1 бесконечный повтор; установить паузу Ctrl.Paused = true
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     1,     // экспоненциальное усиление - норма 2
		Volume:   0.5,   // громкость от 0 до 10
		Silent:   false, // false/true - вкл/выкл звук
	}
	fast := beep.ResampleRatio(4, 1, volume) // воспроизведение со скоростью 1х

	done := make(chan bool)
	speaker.Play(beep.Seq(fast, beep.Callback(func() {
		done <- true
	})))

	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	for {
		select {
		case <-stopPlay:
			speaker.Close()
			return
		case <-done:
			events <- autoNext
			return
		case <-ticker.C:
			speaker.Lock()
			fraction = float32(streamer.Position()) / float32(streamer.Len())
			timePlay = fmt.Sprintf("%s / %s",
				format.SampleRate.D(streamer.Position()).Round(time.Second).String(),
				format.SampleRate.D(streamer.Len()).Round(time.Second).String())
			speaker.Unlock()
		}
	}
}

func stopPlayMp3() {
	for i := 0; i < countPlay; i++ {
		stopPlay <- struct{}{}
	}
}

func OnPlay() {
	events <- play
}

func OnPlayStop() {
	events <- stop
}

func OnPlayNext() {
	events <- next
}

func listDClick(idxPlay int) {
	idx = idxPlay
	events <- dClick
}
