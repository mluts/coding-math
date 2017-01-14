package main

import (
	"../sdl/arrow"
	"../sdl/solsys"
	"../util"
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
	solarSystem  solsys.System
}

var (
	winTitle                       string = "Testing SDL"
	winWidth, winHeight, frameRate int    = 1024, 700, 50
	state                          State
)

func initState() {
	state = State{
		running:      true,
		lastUpdateMs: 0,
		solarSystem: solsys.New(
			float64(winWidth)/2, float64(winHeight)/2,
			30, 20000)}

	for i := 0; i < 20; i++ {
		state.solarSystem.SpawnPlanet(120+float64(i)*rand.Float64()*10, 10, 1, 5+rand.Float64()*10)
	}
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
	var (
		currentTick         = sdl.GetTicks()
		dt          float64 = float64(currentTick-state.lastUpdateMs) / 1000
	)

	for i := 0; i < 30; i++ {
		n := 100
		for i := 0; i < n; i++ {
			state.solarSystem.Update(dt / float64(n))
		}
	}
}

func draw(state *State, renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)
	state.solarSystem.Sun.Draw(renderer, 255, 255, 0, sdl.ALPHA_OPAQUE)

	for i := range state.solarSystem.Planets {
		planet := state.solarSystem.Planets[i]
		state.solarSystem.Planets[i].Draw(renderer, 0+uint8(i)*28, 255-uint8(i)*28, 0, sdl.ALPHA_OPAQUE)

		pos := planet.Particle.Position

		vel := planet.Particle.Velocity
		vel.AddVec(pos)

		acc := planet.Particle.Acceleration
		// fmt.Printf("Len %f Acc %f\n", acc.Len(), (acc.Angle()+math.Pi)*360/(math.Pi*2))
		// acc.Mul(4)
		acc.AddVec(pos)

		renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)

		arrow.NewArrow(
			sdl.Point{
				int32(util.Clampf(pos.X, 0, float64(winWidth))),
				int32(util.Clampf(pos.Y, 0, float64(winHeight)))},
			sdl.Point{
				int32(util.Clampf(vel.X, 0, float64(winWidth))),
				int32(util.Clampf(vel.Y, 0, float64(winHeight)))},
			10).Draw(renderer)

		// renderer.SetDrawColor(255, 0, 0, sdl.ALPHA_OPAQUE)
		//
		// arrow.NewArrow(
		// 	sdl.Point{
		// 		int32(util.Clampf(pos.X, 0, float64(winWidth))),
		// 		int32(util.Clampf(pos.Y, 0, float64(winHeight)))},
		// 	sdl.Point{
		// 		int32(util.Clampf(acc.X, 0, float64(winWidth))),
		// 		int32(util.Clampf(acc.Y, 0, float64(winHeight)))},
		// 	10).Draw(renderer)
	}
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
