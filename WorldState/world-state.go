package state

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	cellWidth  = 10
	cellHeight = 10
)

type Cell struct { // getContent
	x       int32
	y       int32
	content []*string
}

type State struct { // GetContent(X,Y) UpdateContent(X,Y,content)
	cells    []*Cell
	renderer *sdl.Renderer
}

func New(windowHeight, windowWidth int, renderer *sdl.Renderer) *State {
	var cells []*Cell
	var i, j int32
	for i = 0; i < int32(windowWidth); i += cellWidth {
		for j = 0; j < int32(windowHeight); j += cellHeight {
			cells = append(cells, &Cell{x: i, y: j})
		}

	}
	return &State{cells: cells, renderer: renderer}
}

func (s *State) Draw() {
	for _, c := range s.cells {
		c.Draw(s.renderer)
	}
}

func (s *State) UpdateCellContent(x, y int32, value string) {
	x = (x / 10) * 10
	y = (y / 10) * 10
	for _, c := range s.cells {
		if c.x == x && c.y == y {
			c.content = append(c.content, &value)
			return
		}
	}

}

func (c *Cell) Draw(renderer *sdl.Renderer) {
	if len(c.content) == 0 {
		renderer.SetDrawColor(100, 100, 100, 255)
	} else {
		renderer.SetDrawColor(200, 0, 0, 255)
	}

	renderer.DrawRect(&sdl.Rect{X: c.x, Y: c.y, W: cellWidth, H: cellHeight})
}
