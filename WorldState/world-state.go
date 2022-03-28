package state

import (
	. "github.com/fpiwowarczyk/ants-sim/Common"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	cellWidth  = 10
	cellHeight = 10
)

type Cell struct { // getContent
	pos     *Coords
	content []*string
}

type State struct { // GetContent(X,Y) UpdateContent(X,Y,content)
	cells    map[Coords]*Cell
	renderer *sdl.Renderer
}

func New(windowHeight, windowWidth int, renderer *sdl.Renderer) *State {
	cells := make(map[Coords]*Cell)
	var i, j int32
	for i = 0; i < int32(windowWidth); i += cellWidth {
		for j = 0; j < int32(windowHeight); j += cellHeight {
			position := Coords{X: i, Y: j}
			cells[position] = &Cell{pos: &position}
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
	s.cells[Coords{X: x, Y: y}].content = append(s.cells[Coords{X: x, Y: y}].content, &value)
}

func (c *Cell) Draw(renderer *sdl.Renderer) {
	if len(c.content) == 0 {
		renderer.SetDrawColor(100, 100, 100, 255)
	} else {
		renderer.SetDrawColor(200, 0, 0, 255)
	}

	renderer.DrawRect(&sdl.Rect{X: c.pos.X, Y: c.pos.Y, W: cellWidth, H: cellHeight})
}
