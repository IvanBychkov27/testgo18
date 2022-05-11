// https://github.com/faiface/beep/tree/master/examples/tutorial/1-hello-beep
// воспроизводит файл mp3
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	fileName := "cmd/sound/mp3/Tekhnologiya_-_Nazhmi_na_knopku_48044272.mp3"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	//speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	sr := format.SampleRate * 1 // скорость воспроизведения файла (*2 - ускоряет в 2 раза!)
	errInit := speaker.Init(sr, sr.N(time.Second/10))
	if errInit != nil {
		log.Fatal(errInit)
	}

	//loop := beep.Loop(3, streamer)        // перемотка с 3й сек
	//fast := beep.ResampleRatio(4, 5, loop) // перемотка с 4й секунды со скоростью 5х

	Ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false} // -1 бесконечный повтор; установить паузу Ctrl.Paused = true
	volume := &effects.Volume{
		Streamer: Ctrl,
		Base:     2,     // экспоненциальное усиление - норма 2
		Volume:   1,     // громкость от 0 до 10
		Silent:   false, // false/true - вкл/выкл звук
	}
	fast := beep.ResampleRatio(4, 1, volume) // воспроизведение со скоростью 1х

	done := make(chan bool)
	speaker.Play(beep.Seq(fast, beep.Callback(func() {
		done <- true
	})))

	//speaker.Play(beep.Seq(streamer, beep.Callback(func() { // классика
	//	done <- true
	//})))

	//Ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false} // установить паузу
	//speaker.Play(Ctrl)

	chValue := make(chan string)
	go values(chValue)

	for {
		//fmt.Print("Press [ENTER] to pause/resume. ")
		//fmt.Scanln()
		//
		//speaker.Lock()
		//Ctrl.Paused = !Ctrl.Paused
		//speaker.Unlock()

		select {
		case d := <-chValue:
			speaker.Lock()
			if d == "+" {
				volume.Volume += 0.5
			}
			if d == "-" {
				volume.Volume -= 0.5
			}
			if d == "1" {
				Ctrl.Paused = !Ctrl.Paused
			}
			speaker.Unlock()
		case <-ctx.Done():
			log.Println("Done...")
			return
		case <-done:
			log.Println("Done...")
			return
			//case <-time.After(time.Second * 5):
			//	speaker.Lock()
			//	fmt.Print(format.SampleRate.D(streamer.Position()).Round(time.Second), " ")
			//	speaker.Unlock()
		}
	}
}

func values(d chan string) {
	var s string
	fmt.Scanln(&s)
	d <- s
}
