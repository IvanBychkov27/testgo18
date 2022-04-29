package main

import (
	g "github.com/AllenDang/giu"
	"math"
	"os"
)

var (
	linedata               []float64
	lineTicks              []g.PlotTicker
	plotLineTitle          string
	xmin, xmax, ymin, ymax float64
	cond                   = g.ConditionAlways
	checkedAutoSize        = true
)

func exitFunc() {
	os.Exit(0)
}

func sinFunc() {
	plotLineTitle = "sin"
	linedata = nil
	xmin = float64(0)
	xmax = float64(200)
	ymin = float64(-1.2)
	ymax = float64(1.2)
	delta := 0.1
	for x := 0.0; x < 100; x += delta {
		linedata = append(linedata, math.Sin(x))
	}
}

func cosFunc() {
	plotLineTitle = "cos"
	linedata = nil
	xmin = float64(0)
	xmax = float64(200)
	ymin = float64(-1.2)
	ymax = float64(1.2)
	delta := 0.1
	for x := 0.0; x < 100; x += delta {
		linedata = append(linedata, math.Cos(x))
	}
}

func sqrtFunc() {
	plotLineTitle = "sqrt"
	linedata = nil
	xmin = float64(-10)
	xmax = float64(1000)
	ymin = float64(-1)
	ymax = float64(10)
	delta := 0.1
	for x := 0.0; x < 100; x += delta {
		linedata = append(linedata, math.Sqrt(x))
	}
}

func degreeFunc() {
	plotLineTitle = "degree"
	linedata = nil
	xmin = float64(0)
	xmax = float64(10)
	ymin = float64(0)
	ymax = float64(100)
	delta := 1.0
	for x := 0.0; x < 11; x += delta {
		linedata = append(linedata, x*x)
	}
}

func autoSizeFunc() {
	if checkedAutoSize {
		cond = g.ConditionAlways
	} else {
		cond = g.ConditionOnce
	}
}

func loop() {
	autoSizeFunc()
	g.SingleWindowWithMenuBar().Layout(
		g.MenuBar().Layout(
			g.Menu("File").Layout(
				g.Menu("function...").Layout(
					g.MenuItem("sin").OnClick(sinFunc),
					g.MenuItem("cos").OnClick(cosFunc),
					g.MenuItem("sqrt").OnClick(sqrtFunc),
					g.MenuItem("degree").OnClick(degreeFunc),
				),
				g.Menu("options...").Layout(
					g.Checkbox("auto sixe", &checkedAutoSize),
				),
				g.MenuItem("Exit").OnClick(exitFunc),
			),
		),

		g.Plot("Graphic").AxisLimits(xmin, xmax, ymin, ymax, cond).XTicks(lineTicks, false).Plots(
			g.PlotLine(plotLineTitle, linedata),
		),
	)
}

func main() {
	sinFunc()
	wnd := g.NewMasterWindow("Graphic function", 1000, 340, 0).
		RegisterKeyboardShortcuts(
			g.WindowShortcut{
				Key:      g.KeyEscape,
				Modifier: g.ModNone,
				Callback: exitFunc},
		)
	wnd.Run(loop)
}
