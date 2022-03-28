package game

import (
	nest "github.com/fpiwowarczyk/ants-sim/Nest"
	state "github.com/fpiwowarczyk/ants-sim/WorldState"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WinWidth  = 1200
	WinHeight = 1200
)

type Game struct {
	WorldState *state.State
	Nest       *nest.Nest
	Renderer   *sdl.Renderer
}

func Init(renderer *sdl.Renderer) *Game {
	state := state.New(WinHeight, WinWidth, renderer)
	game := &Game{Renderer: renderer, Nest: nest.New(renderer, state), WorldState: state}
	game.Nest.SpawnAnts(1)
	return game
}

func (game *Game) Loop() {
	game.WorldState.Draw()
	game.Nest.Live()
}
