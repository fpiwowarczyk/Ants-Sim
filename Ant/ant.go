package ant

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Ant struct {
	x        int32
	y        int32
	renderer *sdl.Renderer
}

var WinWidth, WinHeight int32 = 1200, 800

func New(renderer *sdl.Renderer) *Ant {
	return &Ant{x: rand.Int31n(WinWidth), y: rand.Int31n(WinHeight), renderer: renderer}
}

func (ant *Ant) Draw() {
	ant.renderer.SetDrawColor(47, 249, 36, 255)
	ant.renderer.FillRect(&sdl.Rect{W: 2, H: 5, X: ant.x, Y: ant.y})
}

func (ant *Ant) MoveRandomly() {
	randomDirection := rand.Intn(7)
	switch randomDirection {
	case 0:
		ant.x = ant.x + 1
	case 1:
		ant.x = ant.x - 1
	case 2:
		ant.y = ant.y + 1
	case 3:
		ant.y = ant.y - 1
	case 4:
		ant.x = ant.x + 1
		ant.y = ant.y + 1
	case 5:
		ant.x = ant.x + 1
		ant.y = ant.y - 1
	case 6:
		ant.x = ant.x - 1
		ant.y = ant.y + 1
	case 7:
		ant.x = ant.x - 1
		ant.y = ant.y - 1
	}
}
