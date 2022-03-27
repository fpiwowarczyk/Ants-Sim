package types

import "github.com/veandco/go-sdl2/sdl"

type Ant struct {
	Pos      *Coords
	dir      int32
	tick     int32
	aperence *sdl.Rect
	renderer *sdl.Renderer
}

type Nest struct {
	pos           *Coords
	occupiedSpace []*Coords
	renderer      *sdl.Renderer
}

type Coords struct {
	X int32
	Y int32
}

type WorldState struct {
}
