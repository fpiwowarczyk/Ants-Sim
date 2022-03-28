package nest

import (
	"math/rand"

	ant "github.com/fpiwowarczyk/ants-sim/Ant"
	. "github.com/fpiwowarczyk/ants-sim/Common"
	state "github.com/fpiwowarczyk/ants-sim/WorldState"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width      = 100
	height     = 50
	doorWidth  = 10
	doorHeight = 20
)

type Nest struct {
	pos        *Coords
	ants       []*ant.Ant
	WorldState *state.State
	renderer   *sdl.Renderer
}

func New(renderer *sdl.Renderer, WorldState *state.State) *Nest {
	return &Nest{
		pos:        &Coords{X: 300, Y: 200},
		WorldState: WorldState,
		renderer:   renderer}
}

func (nest *Nest) Live() {
	for _, a := range nest.ants {
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
	for i := 1; i <= value; i++ {
		newX := nest.pos.X + width - doorWidth + rand.Int31n(10) - 10
		newY := nest.pos.Y + height/2 - doorHeight + 25 + rand.Int31n(10) - 10
		nest.ants = append(nest.ants, ant.New(newX, newY, nest.WorldState, nest.renderer))
	}
}
