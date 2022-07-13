// https://github.com/hajimehoshi/ebiten/blob/main/examples/animation/main.go
// https://habr.com/ru/post/671556/ - Go ebiten: разбираемся с рендерингом и позиционированием текста
/*
   Пример: бегущий человечек
*/

package main

import (
	"bytes"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

var (
	runnerImage *ebiten.Image
)

type Game struct {
	count int
	dx    int
}

func (g *Game) Update() error {
	g.count++
	//if g.count%100 == 0 {
	//	g.dx += 10
	//	fmt.Printf("count:%d  dx:%d\n", g.count, g.dx)
	//}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (g.count / 10) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func openImage(fileName string) ([]byte, error) {
	f, errOpen := os.Open(fileName)
	if errOpen != nil {
		return nil, errOpen
	}
	d, errReadAll := ioutil.ReadAll(f)
	if errReadAll != nil {
		return nil, errReadAll
	}
	return d, nil
}

func main() {
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.

	d, errOpenImage := openImage("cmd/ebiten/animation/runner.png")
	if errOpenImage != nil {
		log.Fatal(errOpenImage)
	}

	img, _, err := image.Decode(bytes.NewReader(d))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")

	errRunGame := ebiten.RunGame(&Game{})
	if errRunGame != nil {
		log.Fatal(errRunGame)
	}
}
