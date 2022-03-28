package main

import (
	"fmt"
	"os"

	game "github.com/fpiwowarczyk/ants-sim/Game"
	"github.com/veandco/go-sdl2/sdl"
)

var Title string = "AntSim"
var WinWidth, WinHeight int32 = 1200, 800

func run() int {
	var Window *sdl.Window
	var renderer *sdl.Renderer

	Window, err := sdl.CreateWindow("Ant-sim",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WinWidth,
		WinHeight,
		sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 1
	}
	defer Window.Destroy()

	renderer, err = sdl.CreateRenderer(Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	running := true

	game := game.Init(renderer)
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		game.Loop()

		renderer.Present()
		sdl.Delay(16)

	}

	return 0
}

func main() {
	os.Exit(run())
}
