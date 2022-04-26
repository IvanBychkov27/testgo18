package main

import "github.com/AllenDang/giu"

var text string

func loop() {
	giu.Window("Window").Layout(
		giu.Align(giu.AlignCenter).To(
			giu.Label("I'm a centered label"),   // метка по центру
			giu.Button("I'm a centered button"), // кнопка по центру
		),

		giu.Align(giu.AlignRight).To(
			giu.Label("I'm a alined to right label"), // метка, выровненная по правому краю
			giu.InputText(&text),                     // поле вывода текста
		),

		giu.Align(giu.AlignRight).To(
			giu.Label("I'm the label"), // метка
			giu.Layout{
				giu.Label("I'm the other label embeded in another layout"), // другая метка, встроенная в другой макет
				giu.Label("I'm the next label"),                            // следующая метка
			},
			giu.Label("I'm the last label"), // последняя метка
		),
		giu.Label("Buttons in row:"), // Кнопки в ряд
		giu.Align(giu.AlignCenter).To(
			giu.Row(
				giu.Button("button 1"),
				giu.Button("button 2"),
			),
		),

		giu.Label("manual alignment"), // ручное выравнивание
		giu.AlignManually(
			giu.AlignCenter,
			giu.Button("I'm button with 100 width").Size(100, 30), // кнопка шириной 100
			100,
			false,
		),
		giu.AlignManually(
			giu.AlignCenter,
			giu.InputText(&text), // поле ввода текста
			100,
			true,
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Alignment demo", 640, 480, 0)
	wnd.Run(loop)
}
