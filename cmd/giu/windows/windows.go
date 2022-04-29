// демонстрация 2х окон
package main

import (
	"github.com/AllenDang/giu"
	"os"
)

func exitFunc() {
	os.Exit(0)
}

func loop() {
	w1 := giu.Window("window 1").Pos(30, 20)
	w2 := giu.Window("window 2").Pos(250, 20)

	w1W, w1H := w1.CurrentSize()
	w1X, w1Y := w1.CurrentPosition()

	w1Layout := giu.Layout{
		giu.Labelf("Focused state %t", w1.HasFocus()),
		giu.Labelf("Position: %d, %d", int(w1X), int(w1Y)),
		giu.Labelf("Size: %d, %d", int(w1W), int(w1H)),
		giu.Button("bring to front window 2").OnClick(w2.BringToFront),
	}
	w2Layout := giu.Layout{
		giu.Labelf("Focused state %t", w2.HasFocus()),
		giu.Button("Exit").OnClick(exitFunc),
	}

	w1.Layout(w1Layout)
	w2.Layout(w2Layout)
}

func main() {
	wnd := giu.NewMasterWindow("windows [DEMO]", 640, 480, 0)
	wnd.Run(loop)
}
