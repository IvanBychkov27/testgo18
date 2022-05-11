package main

import (
	"context"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Node struct {
	Files []string
}

func NewNode() *Node {
	return &Node{[]string{}}
}

func main() {
	rand.Seed(time.Now().UnixMilli())

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	wg := &sync.WaitGroup{}

	listMp3 := NewNode()
	listDir(listMp3, "/home/ivan/Iv/mp3", ".mp3")

	l := len(listMp3.Files)
	log.Println("all mp3 files:", l)

	next := make(chan bool)
	var fileName string
	for {
		fileName = listMp3.Files[rand.Intn(l)] // выбор песни из списка
		log.Println(fileName)

		wg.Add(1)
		go playMp3(ctx, wg, fileName, next)

		select {
		case <-next:
		case <-ctx.Done():
			wg.Wait()
			log.Println("Done...")
			return
		}
	}
}

func playMp3(ctx context.Context, wg *sync.WaitGroup, fileName string, next chan bool) {
	defer wg.Done()

	file, errOpen := os.Open(fileName)
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
		case <-ctx.Done():
			return
		case <-done:
			next <- true
			return
		}
	}
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
