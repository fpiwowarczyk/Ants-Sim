package main

import (
	"fmt"
	"os"

	nest "github.com/fpiwowarczyk/ants-sim/Nest"
	"github.com/veandco/go-sdl2/sdl"
)

var Title string = "AntSim"
var WinWidth, WinHeight int32 = 1200, 800

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	window, err := sdl.CreateWindow("Ant-sim",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WinWidth,
		WinHeight,
		sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	running := true

	nest := nest.New(renderer)
	nest.SpawnAnts(10)

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		nest.Live()

		renderer.Present()
		sdl.Delay(16)

	}

	return 0
}

func main() {
	os.Exit(run())
}
