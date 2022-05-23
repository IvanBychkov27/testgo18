package main

import (
	g "github.com/AllenDang/giu"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

var (
	buttonPlayDisabled = false
	buttonStopDisabled = true
	stopPlay           chan struct{}
	pause              chan struct{}
	valueChange        chan struct{}
	value              int32 = 30
	labelInfo          string
)

func main() {
	stopPlay = make(chan struct{})
	valueChange = make(chan struct{})
	pause = make(chan struct{})

	wnd := g.NewMasterWindow("mp3", 400, 300, 0).
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
			g.Button("Play").OnClick(OnPlay).Disabled(buttonPlayDisabled),
			g.Button("Pause").OnClick(OnPause).Disabled(buttonStopDisabled),
			g.Button("Stop").OnClick(OnPlayStop).Disabled(buttonStopDisabled),
		),
		g.Label(labelInfo),
		g.Row(
			g.Label("Value"),
			g.SliderInt(&value, 0, 100).Size(100).OnChange(OnValue),
		),
	)
}

func exitFunc() {
	os.Exit(0)
}

func OnPlay() {
	buttonPlayDisabled = true
	buttonStopDisabled = false

	labelInfo = "play"
	go playMp3()
}

func OnPlayStop() {
	stopPlay <- struct{}{}
}

func OnPause() {
	pause <- struct{}{}
}

func OnValue() {
	if buttonPlayDisabled {
		valueChange <- struct{}{}
	}
}

func playMp3() {
	file, errOpen := os.Open("cmd/giu_example/mp3/Ruki_Vverkh_-_On_tebya_celuet.mp3")
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
		Base:     float64(value) / 100, // экспоненциальное усиление - норма 2
		Volume:   1,                    // громкость от 0 до 10
		Silent:   false,                // false/true - вкл/выкл звук
	}
	fast := beep.ResampleRatio(4, 1, volume) // воспроизведение со скоростью 1х

	done := make(chan bool)
	speaker.Play(beep.Seq(fast, beep.Callback(func() {
		done <- true
	})))
	defer speaker.Close()

	for {
		select {
		case <-stopPlay:
			buttonEnabled()
			return
		case <-done:
			buttonEnabled()
			return

		case <-pause:
			speaker.Lock()
			ctrl.Paused = !ctrl.Paused
			if ctrl.Paused {
				labelInfo = "pause"
			} else {
				labelInfo = "play"
			}

			speaker.Unlock()

		case <-valueChange:
			speaker.Lock()
			volume.Base = float64(value) / 100 // громкость
			speaker.Unlock()
		}
	}
}

func buttonEnabled() {
	buttonPlayDisabled = false
	buttonStopDisabled = true
	labelInfo = ""
}
