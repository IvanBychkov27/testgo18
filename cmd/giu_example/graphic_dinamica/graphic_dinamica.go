package main

import (
	g "github.com/AllenDang/giu"
	"math"
	"os"
	"time"
)

var (
	linedataSin            []float64
	linedataCos            []float64
	lineTicks              []g.PlotTicker
	plotLineTitleSin       string
	plotLineTitleCos       string
	cond                   = g.ConditionAlways
	checkedAutoSize        = true
	checkedSin, checkedCos bool
	xmin                   = -3.0
	xmax                   = 330.0
	ymin                   = -1.2
	ymax                   = 1.2
	delta                  = 0.1

	checkboxFuncEnabled = true
)

func exitFunc() {
	os.Exit(0)
}

func sinFunc() {
	checkboxFuncEnabled = false
	plotLineTitleSin = "sin"
	for x := 0.0; x < xmax*delta; x += delta {
		linedataSin = append(linedataSin, math.Sin(x))
		time.Sleep(time.Millisecond * 10)
	}
	checkboxFuncEnabled = true
}

func cosFunc() {
	checkboxFuncEnabled = false
	plotLineTitleCos = "cos"
	for x := 0.0; x < xmax*delta; x += delta {
		linedataCos = append(linedataCos, math.Cos(x))
		time.Sleep(time.Millisecond * 10)
	}
	checkboxFuncEnabled = true
}

func autoSizeFunc() {
	if checkedAutoSize {
		cond = g.ConditionAlways
	} else {
		cond = g.ConditionOnce
	}
}

func showSin() {
	if checkedSin {
		go sinFunc()
	} else {
		linedataSin = nil
	}
}

func showCos() {
	if checkedCos {
		go cosFunc()
	} else {
		linedataCos = nil
	}
}

func loop() {
	autoSizeFunc()
	g.SingleWindowWithMenuBar().Layout(
		g.MenuBar().Layout(
			g.Menu("File").Layout(
				g.Menu("function...").Layout(
					g.Checkbox("sin", &checkedSin).OnChange(showSin),
					g.Checkbox("cos", &checkedCos).OnChange(showCos),
				).Enabled(checkboxFuncEnabled),
				g.Menu("options...").Layout(
					g.Checkbox("auto sixe", &checkedAutoSize),
				),
				g.MenuItem("Exit").OnClick(exitFunc),
			),
		),

		g.Plot("Graphic").AxisLimits(xmin, xmax, ymin, ymax, cond).XTicks(lineTicks, false).Plots(
			g.PlotLine(plotLineTitleSin, linedataSin),
			g.PlotLine(plotLineTitleCos, linedataCos),
		),
	)
}

func main() {
	wnd := g.NewMasterWindow("Graphic function", 1000, 340, 0).
		RegisterKeyboardShortcuts(
			g.WindowShortcut{
				Key:      g.KeyEscape,
				Modifier: g.ModNone,
				Callback: exitFunc},
		)
	wnd.Run(loop)
}
