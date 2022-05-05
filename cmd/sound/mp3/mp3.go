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

	file, err := os.Open("cmd/sound/mp3/Tekhnologiya_-_Nazhmi_na_knopku_48044272.mp3")
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

	ctrl := &beep.Ctrl{Streamer: beep.Loop(1, streamer), Paused: false} // -1 бесконечный повтор; установить паузу ctrl.Paused = true
	volume := &effects.Volume{
		Streamer: ctrl,
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

	//ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false} // установить паузу
	//speaker.Play(ctrl)

	for {

		//fmt.Print("Press [ENTER] to pause/resume. ")
		//fmt.Scanln()
		//
		//speaker.Lock()
		//ctrl.Paused = !ctrl.Paused
		//speaker.Unlock()

		select {
		case <-ctx.Done():
			log.Println("Done...")
			return
		case <-done:
			log.Println("Done...")
			return
		case <-time.After(time.Second * 5):
			speaker.Lock()
			fmt.Print(format.SampleRate.D(streamer.Position()).Round(time.Second), " ")
			speaker.Unlock()
		}
	}
}
