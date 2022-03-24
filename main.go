package main

import (
	"fmt"
	"os"

	ant "github.com/fpiwowarczyk/ants-sim/Ant"
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

	var ants []*ant.Ant

	ants = append(ants, ant.New(renderer))
	period := 0
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		for _, a := range ants {
			a.Draw()
			a.MoveRandomly()
		}
		if period == 1 {
			fmt.Println("Ant added", len(ants))
			ants = append(ants, ant.New(renderer))
			period = 0
		}
		period++

		renderer.Present()
		sdl.Delay(16)

	}
	return 0
}

func main() {
	fmt.Println("Run")
	os.Exit(run())
}
