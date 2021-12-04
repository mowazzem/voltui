package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	volgo "github.com/itchyny/volume-go"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	vol, err := volgo.GetVolume()
	if err != nil {
		panic("error")
	}

	p := helpDialog()
	g0 := volumeBarDialog(vol)

	draw(p, g0)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		eid := e.ID
		switch eid {
		case "q", "<C-c>":
			return
		case "<Right>":
			{
				if vol <= 95 {
					volgo.IncreaseVolume(5)
					vol += 5
					g0.Percent = vol
					draw(p, g0)

				}
			}
		case "<Left>":
			{
				if vol >= 5 {
					volgo.IncreaseVolume(-5)
					vol -= 5
					g0.Percent = vol
					draw(p, g0)

				}
			}
		}
	}
}

func helpDialog() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = `1. press 'q' for quit, 2. press Right for increase, 3. press Left for decrease`
	p.SetRect(0, 0, 90, 3)
	return p
}

func volumeBarDialog(vol int) *widgets.Gauge {
	g0 := widgets.NewGauge()
	g0.Title = "Volume"
	g0.SetRect(0, 3, 90, 7)
	g0.Percent = vol
	g0.BarColor = ui.ColorGreen
	g0.BorderStyle.Fg = ui.ColorRed
	g0.TitleStyle.Fg = ui.ColorYellow
	return g0
}

func draw(drawables ...ui.Drawable) {
	ui.Clear()
	ui.Render(drawables...)
}
