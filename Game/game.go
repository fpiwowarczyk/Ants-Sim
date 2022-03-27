package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	Renderer *sdl.Renderer
	State    *WorldState
}

func New(renderer *sdl.Renderer) *Game {
	game := &Game{}
	game.state.Nest = nest.New(renderer)
	game.Renderer = renderer
	game.Nest.SpawnAnts(10)

	return game
}

func (game *Game) Loop() {
	game.Nest.Live()
}
