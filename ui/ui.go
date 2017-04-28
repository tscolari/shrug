package ui

import (
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
	"github.com/tscolari/shrug/garden"
)

var app = &views.Application{}
var window = &mainWindow{}

type GardenClient interface {
	Containers() []garden.Container
}

func Start(client GardenClient) error {
	containers := client.Containers()
	window.model = newContainersModel(containers)
	window.client = client

	title := views.NewTextBar()
	title.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorTeal).
		Foreground(tcell.ColorWhite))
	title.SetCenter("SHRUG", tcell.StyleDefault)
	title.SetRight("Example v1.0", tcell.StyleDefault)

	window.keybar = views.NewSimpleStyledText()
	window.keybar.RegisterStyle('N', tcell.StyleDefault.
		Background(tcell.ColorSilver).
		Foreground(tcell.ColorBlack))
	window.keybar.RegisterStyle('A', tcell.StyleDefault.
		Background(tcell.ColorSilver).
		Foreground(tcell.ColorRed))
	window.keybar.SetMarkup("[%AQ%N] Quit")

	window.status = views.NewSimpleStyledTextBar()
	window.status.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlue).
		Foreground(tcell.ColorYellow))
	window.status.RegisterLeftStyle('N', tcell.StyleDefault.
		Background(tcell.ColorYellow).
		Foreground(tcell.ColorBlack))

	window.status.SetLeft("My status is here.")
	window.status.SetRight("%UCellView%N demo!")
	window.status.SetCenter("Cen%ST%Ner")

	window.main = views.NewCellView()
	window.main.SetModel(window.model)
	window.main.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlack))

	window.SetMenu(window.keybar)
	window.SetTitle(title)
	window.SetContent(window.main)
	window.SetStatus(window.status)

	window.updateKeys()

	app.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	app.SetRootWidget(window)

	go func(client GardenClient) {
		for {
			window.model.Refresh(client)
			time.Sleep(time.Second)
		}
	}(client)

	if err := app.Run(); err != nil {
		return err
	}

	return nil
}
