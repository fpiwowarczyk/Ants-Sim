package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	window, err := sdl.CreateWindow("winTitle",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800,
		800,
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
}

func main() {
	fmt.Println("Run")
	os.Exit(run())
}
