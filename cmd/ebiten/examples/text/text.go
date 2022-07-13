// https://habr.com/ru/post/671556/
/*
	Выравнивание текста
*/
package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

const (
	windowWidth  = 32 * 14
	windowHeight = 32 * 8
)

type Game struct{}

func (g *Game) Update() error { return nil }

func (g *Game) Draw(screen *ebiten.Image) {
	// Локальные сокращения, чтобы уменьшить код по ширине
	const w = windowWidth
	const h = windowHeight

	ebitenutil.DrawRect(screen, 0, 0, w, h, color.White)

	// Рисуем сетку (32x32 и 64x64)
	gridColor64 := &color.RGBA{A: 50}
	gridColor32 := &color.RGBA{A: 20}
	for y := 0.0; y < h; y += 32 {
		ebitenutil.DrawLine(screen, 0, y, w, y, gridColor32)
	}
	for y := 0.0; y < h; y += 64 {
		ebitenutil.DrawLine(screen, 0, y, w, y, gridColor64)
	}
	for x := 0.0; x < w; x += 32 {
		ebitenutil.DrawLine(screen, x, 0, x, h, gridColor32)
	}
	for x := 0.0; x < w; x += 64 {
		ebitenutil.DrawLine(screen, x, 0, x, h, gridColor64)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func main() {
	//ctx := newContext()
	//game := &Game{ctx: ctx}
	ebiten.SetWindowTitle("Text Rendering")
	ebiten.SetWindowSize(windowWidth, windowHeight)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
