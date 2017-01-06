package main

import (
	"../sdl/arrow"
	"../sdl/point"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	runtime.LockOSThread()
	rand.Seed(time.Now().Unix())
}

type State struct {
	running        bool
	lastUpdateMs   uint32
	speedMs        int
	lastTick       uint32
	mouseX, mouseY int32
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
		}
	case *sdl.MouseMotionEvent:
		state.mouseX = t.X
		state.mouseY = t.Y
	}
}

func updateState(state *State) {
	x, y, _ := sdl.GetMouseState()
	state.mouseX = int32(x)
	state.mouseY = int32(y)
}

func draw(state *State, renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)

	p := sdl.Point{int32(winWidth / 2), int32(winHeight / 2)}
	m := sdl.Point{state.mouseX, state.mouseY}
	p = point.Translate(p, -m.X, -m.Y)
	angle := math.Atan2(float64(p.Y), float64(p.X)) + math.Pi
	p = point.Translate(p, m.X, m.Y)

	a := arrow.NewArrow(sdl.Point{p.X - 50, p.Y}, sdl.Point{p.X + 50, p.Y}, 50)
	a = a.Rotate(p, angle)
	for _, l := range a.Lines() {
		renderer.DrawLines([]sdl.Point{l.From, l.To})
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

	frameRate := gfx.FPSmanager{}
	gfx.InitFramerate(&frameRate)
	gfx.SetFramerate(&frameRate, 50)

	state = State{
		running:      true,
		speedMs:      0,
		lastUpdateMs: 0,
		mouseX:       0,
		mouseY:       0}

	for state.running {
		gfx.FramerateDelay(&frameRate)

		processInput(&state, sdl.PollEvent())
		currentTick := sdl.GetTicks()
		if (currentTick - state.lastUpdateMs) > uint32(state.speedMs) {
			updateState(&state)
			state.lastUpdateMs = currentTick
		}
		draw(&state, renderer)
		renderer.Present()
	}
}
