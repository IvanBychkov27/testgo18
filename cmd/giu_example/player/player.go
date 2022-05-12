package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	allFiles []string
	nameFile string
	events   chan int
	next     chan bool
	idx      int
)

func main() {
	rand.Seed(time.Now().UnixMilli())

	allFiles = make([]string, 0, 1000)
	events = make(chan int)
	next = make(chan bool)

	path := "/home/ivan/Iv/mp3"
	fileExtension := ".mp3"
	listDir(path, fileExtension)

	go processing()

	wnd := g.NewMasterWindow("Player Iv", 800, 640, 0).
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
				g.MenuItem("Exit").OnClick(exitFunc),
			),
		),
		g.Row(
			g.Label("Play:"),
			g.Label(nameFile),
		),
		g.Row(
			g.Button("Play").OnClick(OnPlay),
			g.Button("Stop").OnClick(OnPlayStop).Disabled(true),
			g.Button("Next").OnClick(OnPlayNext),
			g.Button("Random").OnClick(OnPlayRandom),
		),
		g.ListBox("ListBox1", allFiles).OnDClick(listDClick),
	)
}

func exitFunc() {
	os.Exit(0)
}

func listDClick(idxPlay int) {
	idx = idxPlay
	events <- 5
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

func processing() {
	for {
		select {
		case <-next:
			idx++
		case event := <-events:
			switch event {
			case 1: // play
			//case 2: // stop
			case 3: // next
				idx++
			case 4: // random
				idx = rand.Intn(len(allFiles))
			case 5: // dClick
			}
		}

		nameFile = fmt.Sprintf(" %d - %s", idx, allFiles[idx])
		go playMp3()
	}
}

func playMp3() {
	if idx > len(allFiles)-1 {
		idx = 0
	}

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

	for {
		select {
		//case <-stop:
		case <-done:
			next <- true
			return
		}
	}
}

func OnPlay() {
	events <- 1
}

func OnPlayStop() {
	events <- 2
}

func OnPlayNext() {
	events <- 3
}

func OnPlayRandom() {
	events <- 4
}
