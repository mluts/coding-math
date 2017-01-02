package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	runtime.LockOSThread()
	rand.Seed(time.Now().Unix())
}

type State struct {
	running      bool
	lines        []Line
	lastUpdateMs uint32
	speedMs      int
	lastTick     uint32
}

type Line struct {
	x0, y0, x1, y1 int
}

var (
	winTitle                       string = "Testing SDL"
	winWidth, winHeight, frameRate int    = 800, 600, 50
)

func processInput(state *State, event sdl.Event) {
	switch t := event.(type) {
	case *sdl.QuitEvent:
		state.running = false
	case *sdl.KeyDownEvent:
		switch t.Keysym.Sym {
		case sdl.K_ESCAPE:
			state.running = false
		case sdl.K_UP:
			state.speedMs += 100
		case sdl.K_DOWN:
			state.speedMs -= 100
			if state.speedMs < 0 {
				state.speedMs = 0
			}
		}
	}
}

func ensureFrameRate(state *State) {
	currentTick := sdl.GetTicks()
	sleepMs := int(1000/frameRate - int(currentTick-state.lastTick))
	if sleepMs > 0 {
		sdl.Delay(uint32(sleepMs))
	}
	state.lastTick = currentTick
}

func updateState(state *State) {
	currentTick := sdl.GetTicks()

	if (currentTick - state.lastUpdateMs) > uint32(state.speedMs) {
		state.lastUpdateMs = currentTick
		line := Line{rand.Intn(winWidth), rand.Intn(winHeight),
			rand.Intn(winWidth), rand.Intn(winHeight)}

		state.lines = append(state.lines, line)
	}
}

func draw(state *State, renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)
	for _, line := range state.lines {
		renderer.DrawLine(line.x0, line.y0, line.x1, line.y1)
	}
}

func main() {
	var (
		window   *sdl.Window
		renderer *sdl.Renderer
		err      error
		state    State
	)
	sdl.Init(sdl.INIT_VIDEO)

	window, err = sdl.CreateWindow(
		winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic("Can't initialize SDL window")
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(
		window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic("Can't initialize SDL renderer")
	}
	defer renderer.Destroy()
	defer sdl.Quit()

	state = State{
		running:      true,
		speedMs:      1000,
		lastUpdateMs: 0,
		lastTick:     0}

	for state.running {
		ensureFrameRate(&state)
		processInput(&state, sdl.PollEvent())
		updateState(&state)
		draw(&state, renderer)
		renderer.Present()
	}
}
