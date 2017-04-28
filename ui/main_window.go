package ui

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

type mainWindow struct {
	client GardenClient
	view   views.View
	main   *views.CellView
	keybar *views.SimpleStyledText
	status *views.SimpleStyledTextBar
	model  *containersModel

	views.Panel
}

func (a *mainWindow) HandleEvent(ev tcell.Event) bool {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyCtrlL:
			app.Refresh()
			return true
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'Q', 'q':
				app.Quit()
				return true
			case 'S', 's':
				a.model.hide = false
				a.updateKeys()
				return true
			case 'H', 'h':
				a.model.hide = true
				a.updateKeys()
				return true
			case 'E', 'e':
				a.model.enab = true
				a.updateKeys()
				return true
			case 'D', 'd':
				a.model.enab = false
				a.updateKeys()
				return true
			}
		}
	}
	return a.Panel.HandleEvent(ev)
}

func (a *mainWindow) Draw() {
	a.status.SetLeft(a.model.loc)
	a.Panel.Draw()
}

func (a *mainWindow) updateKeys() {
	m := a.model
	w := "[%AQ%N] Quit"
	if !m.enab {
		w += "  [%AE%N] Enable cursor"
	} else {
		w += "  [%AD%N] Disable cursor"
		if !m.hide {
			w += "  [%AH%N] Hide cursor"
		} else {
			w += "  [%AS%N] Show cursor"
		}
	}
	a.keybar.SetMarkup(w)
	app.Update()
}
