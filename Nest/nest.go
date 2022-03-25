package nest

import (
	"math/rand"

	ant "github.com/fpiwowarczyk/ants-sim/Ant"
	"github.com/veandco/go-sdl2/sdl"
)

type Nest struct {
	x        int32
	y        int32
	ants     []*ant.Ant
	renderer *sdl.Renderer
}

const (
	width      = 100
	height     = 50
	doorWidth  = 10
	doorHeight = 20
)

func New(renderer *sdl.Renderer) *Nest {
	return &Nest{x: 300,
		y:        200,
		renderer: renderer}
}

func (nest *Nest) Live() { //Maybe i should use textuer here
	for _, a := range nest.ants {
		a.Draw()
		a.Move()
	}
	nest.draw()
}

func (nest *Nest) draw() {
	nest.renderer.SetDrawColor(0, 105, 148, 255)
	nest.renderer.FillRect(&sdl.Rect{X: nest.x, Y: nest.y, W: width, H: height})
	nest.renderer.SetDrawColor(106, 67, 33, 255)
	nest.renderer.FillRect(&sdl.Rect{X: nest.x + width - doorWidth, Y: nest.y + height/2 - doorHeight + 10, W: doorWidth, H: doorHeight})
}

func (nest *Nest) SpawnAnts(value int) {
	for i := 0; i <= value; i++ {
		nest.ants = append(nest.ants, ant.New(nest.x+width-doorWidth+rand.Int31n(10)-10, nest.y+height/2-doorHeight+25+rand.Int31n(10)-10, nest.renderer))
	}
}
