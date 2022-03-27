package game

import (
	. "github.com/fpiwowarczyk/ants-sim/Game/Common"
	"github.com/veandco/go-sdl2/sdl"
)

type Nest struct {
	pos           *Coords
	occupiedSpace []*Coords
	renderer      *sdl.Renderer
}

type WorldState struct {
}
