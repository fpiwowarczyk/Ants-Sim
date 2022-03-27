package nest

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	width      = 100
	height     = 50
	doorWidth  = 10
	doorHeight = 20
)

func New(renderer *sdl.Renderer) *Nest {
	return &Nest{
		pos:      &Coords{X: 300, Y: 200},
		renderer: renderer}
}

func (nest *Nest) Live() {
	for _, a := range game.Ants {
		a.Draw()
		a.Move()
	}
	nest.draw()
}

func (nest *Nest) draw() {
	nest.renderer.SetDrawColor(0, 105, 148, 255)
	nest.renderer.FillRect(&sdl.Rect{X: nest.pos.X, Y: nest.pos.Y, W: width, H: height})
	nest.renderer.SetDrawColor(106, 67, 33, 255)
	nest.renderer.FillRect(&sdl.Rect{X: nest.pos.X + width - doorWidth, Y: nest.pos.Y + height/2 - doorHeight + 10, W: doorWidth, H: doorHeight})
}

func (nest *Nest) SpawnAnts(value int) {
	for i := 0; i <= value; i++ {
		nest.Ants = append(nest.Ants, ant.New(nest.pos.X+width-doorWidth+rand.Int31n(10)-10, nest.pos.Y+height/2-doorHeight+25+rand.Int31n(10)-10, nest.renderer))
	}
}
