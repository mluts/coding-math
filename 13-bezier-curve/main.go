package main

import (
	"../sdl/bezier"
	"../sdl/vector"
	// "../sdl/arrow"
	// "../sdl/solsys"
	// "../util"
	// "fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
	// "math"
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
	lastUpdateMs uint32
}

var (
	winTitle                       string = "Testing SDL"
	winWidth, winHeight, frameRate int    = 1024, 700, 50
	state                          State
)

func initState() {
	state = State{
		running:      true,
		lastUpdateMs: 0}
}

func processInput(state *State, event sdl.Event) {
	switch t := event.(type) {
	case *sdl.QuitEvent:
		state.running = false
	case *sdl.KeyDownEvent:
		switch t.Keysym.Sym {
		case sdl.K_ESCAPE:
			state.running = false
		}
	}
}

func updateState(state *State) {
	// var (
	// 	currentTick         = sdl.GetTicks()
	// 	dt          float64 = float64(currentTick-state.lastUpdateMs) / 1000
	// )
}

func draw(state *State, renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)

	points := make([]sdl.Point, 0)
	p0 := vec.NewVec2(100, 500)
	p1 := vec.NewVec2(300, 100)
	p2 := vec.NewVec2(500, 300)

	for _, p := range []vec.Vec2{p0, p1, p2} {
		gfx.FilledCircleRGBA(renderer, int(p.X), int(p.Y), 3, 0, 0, 0, sdl.ALPHA_OPAQUE)
	}

	renderer.DrawLines([]sdl.Point{
		sdl.Point{int32(p0.X), int32(p0.Y)},
		sdl.Point{int32(p1.X), int32(p1.Y)}})

	renderer.DrawLines([]sdl.Point{
		sdl.Point{int32(p1.X), int32(p1.Y)},
		sdl.Point{int32(p2.X), int32(p2.Y)}})

	for t := float64(0); t <= 1; t += 0.1 {
		b0 := bez.LinearBezier(p0, p1, t)
		b1 := bez.LinearBezier(p1, p2, t)
		renderer.DrawLines([]sdl.Point{
			sdl.Point{int32(b0.X), int32(b0.Y)},
			sdl.Point{int32(b1.X), int32(b1.Y)}})

		b := bez.QuadraticBezier(p0, p1, p2, t)
		points = append(points, sdl.Point{int32(b.X), int32(b.Y)})
	}

	renderer.DrawLines(points)
}

func main() {
	var (
		window   *sdl.Window
		renderer *sdl.Renderer
		err      error
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
	gfx.SetFramerate(&frameRate, 100)

	initState()

	for state.running {
		gfx.FramerateDelay(&frameRate)

		processInput(&state, sdl.PollEvent())
		currentTick := sdl.GetTicks()
		updateState(&state)
		state.lastUpdateMs = currentTick
		draw(&state, renderer)
		renderer.Present()
	}
}
