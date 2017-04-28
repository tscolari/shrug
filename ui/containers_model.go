package ui

import (
	"fmt"
	"sync"

	"github.com/gdamore/tcell"
	"github.com/tscolari/shrug/garden"
)

func newContainersModel(containers []garden.Container) *containersModel {
	return &containersModel{
		containers: containers,
		endx:       1,
		lock:       new(sync.Mutex),
	}
}

type containersModel struct {
	containers []garden.Container
	lock       *sync.Mutex

	x    int
	y    int
	endx int
	hide bool
	enab bool
	loc  string
}

func (m *containersModel) Refresh(client GardenClient) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.containers = client.Containers()
}

func (m *containersModel) GetBounds() (int, int) {
	m.lock.Lock()
	defer m.lock.Unlock()

	return len(m.containers[0].Name), len(m.containers)
}

func (m *containersModel) MoveCursor(offx, offy int) {
	m.x += offx
	m.y += offy
	m.limitCursor()
}

func (m *containersModel) limitCursor() {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.x < 0 {
		m.x = 0
	}
	if m.x > m.endx-1 {
		m.x = m.endx - 1
	}
	if m.y < 0 {
		m.y = 0
	}
	if m.y > len(m.containers)-1 {
		m.y = len(m.containers) - 1
	}
	m.loc = fmt.Sprintf("Cursor is %d,%d", m.x, m.y)
}

func (m *containersModel) GetCursor() (int, int, bool, bool) {
	return m.x, m.y, m.enab, !m.hide
}

func (m *containersModel) SetCursor(x int, y int) {
	m.x = x
	m.y = y

	m.limitCursor()
}

func (m *containersModel) GetCell(x, y int) (rune, tcell.Style, []rune, int) {
	m.lock.Lock()
	defer m.lock.Unlock()

	style := tcell.StyleDefault

	if len(m.containers) <= y {
		return ' ', style, nil, 1
	}

	container := m.containers[y]

	if len(container.Name)+2 == x {
		style := style.
			Foreground(tcell.ColorBlack).
			Background(tcell.ColorWhite)
		return '|', style, nil, 1
	}

	if len(container.Name) <= x {
		return ' ', style, nil, 1
	}

	ch := rune(m.containers[y].Name[x])
	style = style.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack)
	return ch, style, nil, 1
}
