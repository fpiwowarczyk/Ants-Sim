package ant

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Ant struct {
	Pos      *Coords
	dir      int32
	tick     int32
	aperence *sdl.Rect
	renderer *sdl.Renderer
}

var WinWidth, WinHeight int32 = 1200, 800

func New(x, y int32, renderer *sdl.Renderer) *Ant {
	return &Ant{
		Pos:      &Coords{X: x, Y: y},
		dir:      rand.Int31n(4),
		tick:     rand.Int31n(10),
		renderer: renderer}
}

func (ant *Ant) Draw() {
	ant.tick++
	ant.renderer.SetDrawColor(47, 249, 36, 255)
	ant.renderer.FillRect(&sdl.Rect{W: 2, H: 5, X: ant.Pos.X, Y: ant.Pos.Y})
}

func (ant *Ant) Move() {
	if ant.tick < 25 {
		ant.dir = 1
	} else {
		if ant.tick%5 == 0 {
			newDirection := rand.Int31n(3) - 1
			ant.dir += newDirection
			if ant.dir <= -1 {
				ant.dir = 3
			} else if ant.dir >= 4 {
				ant.dir = 0
			}

		}
	}

	switch ant.dir {
	case 0:
		ant.Pos.Y -= 1
	case 1:
		ant.Pos.X += 1
	case 2:
		ant.Pos.Y += 1
	case 3:
		ant.Pos.X -= 1
	}

	if ant.Pos.X >= 1200 {
		ant.Pos.X = -5
	} else if ant.Pos.X <= -10 {
		ant.Pos.X = 1205
	} else if ant.Pos.Y <= -10 {
		ant.Pos.Y = 805
	} else if ant.Pos.Y >= 810 {
		ant.Pos.Y = -5
	}

}

func borderSceen() {

}
