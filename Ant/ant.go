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

}
