// https://golangrepo.com/repo/fogleman-gg-go-images-tools
// примеры: https://github.com/fogleman/gg/tree/master/examples
// скачать шрифты см: https://andryushkin.ru/fonts/download-font-arial/?ysclid=l6nmzur8zb701862790

package main

import (
	"github.com/fogleman/gg"
	"log"
	"math"
)

type Point struct {
	X, Y float64
}

func Polygon(n int, x, y, r float64) []Point {
	result := make([]Point, n)
	for i := 0; i < n; i++ {
		a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
		result[i] = Point{x + r*math.Cos(a), y + r*math.Sin(a)}
	}
	return result
}

func main() {
	//circleDraw("cmd/graphics/pic/circle.png")
	//starDraw("cmd/graphics/pic/star.png")
	textDraw("cmd/graphics/pic/text.png", "Hello, world!")

	log.Println("Done..")
}

func textDraw(fileName, text string) {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	//font := "cmd/graphics/fonts/Arial/ArialBlack/ArialBlack.ttf"
	//font := "cmd/graphics/fonts/Arial/ArialBold/ArialBold.ttf"
	//font := "cmd/graphics/fonts/Arial/ArialItalic/ArialItalic.ttf"
	font := "cmd/graphics/fonts/Arial/ArialBoldItalic/ArialBoldItalic.ttf"
	//font := "cmd/graphics/fonts/Arial/ArialRegular/ArialRegular.ttf"
	err := dc.LoadFontFace(font, 96)
	if err != nil {
		log.Println("error:", err.Error())
	}
	dc.DrawStringAnchored(text, S/2, S/2, 0.5, 0.5)
	dc.SavePNG(fileName)
}

func starDraw(fileName string) {
	n := 5
	points := Polygon(n, 512, 512, 400)
	dc := gg.NewContext(1024, 1024)
	dc.SetHexColor("fff")
	dc.Clear()
	for i := 0; i < n+1; i++ {
		index := (i * 2) % n
		p := points[index]
		dc.LineTo(p.X, p.Y)
	}
	dc.SetRGBA(0, 0.5, 0, 1)
	dc.SetFillRule(gg.FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetRGBA(0, 1, 0, 0.5)
	dc.SetLineWidth(16)
	dc.Stroke()
	dc.SavePNG(fileName)
}

func circleDraw(fileName string) {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG(fileName)
}
