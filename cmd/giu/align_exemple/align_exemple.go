package main

import (
	"github.com/AllenDang/giu"
	"strconv"
	"strings"
)

var x, y, res string

func main() {
	wnd := giu.NewMasterWindow("Alignment exemple", 800, 600, 0)
	wnd.Run(loop)
}

func loop() {
	res = calculation(x, y)

	window := giu.Window("Calculation").Size(400, 200)

	layout := giu.Layout{
		giu.Row(
			giu.Label("x ="),
			giu.AlignManually(
				giu.AlignCenter,
				giu.InputText(&x),
				200,
				true,
			),
		),

		giu.Row(
			giu.Label("y ="),
			giu.AlignManually(
				giu.AlignCenter,
				giu.InputText(&y),
				200,
				true,
			),
		),
		giu.Row(
			giu.Align(giu.AlignCenter).To(
				giu.InputText(&res),
			),
		),
	}

	window.Layout(layout)

}

func calculation(a, b string) string {
	if a == "" || b == "" {
		return ""
	}

	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	xi, errX := strconv.Atoi(a)
	if errX != nil {
		a = "err"
	}
	yi, errY := strconv.Atoi(b)
	if errY != nil {
		b = "err"
	}
	r := "error"
	if errX == nil && errY == nil {
		r = strconv.Itoa(xi + yi)
	}
	return a + " + " + b + " = " + r
}
