package main

import (
	"../sdl/solsys"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
	// "math"
	"../sdl/arrow"
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
	speedMs      int
	solarSystem  solsys.System
}

var (
	winTitle                       string = "Testing SDL"
	winWidth, winHeight, frameRate int    = 1024, 700, 50
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
	}
}

func updateState(state *State) {
	var (
		currentTick         = sdl.GetTicks()
		dt          float64 = float64(currentTick-state.lastUpdateMs) / 1000
	)

	state.solarSystem.Update(dt)
}

func draw(state *State, renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)
	state.solarSystem.Sun.Draw(renderer, 255, 255, 0, sdl.ALPHA_OPAQUE)

	for i := range state.solarSystem.Planets {
		planet := state.solarSystem.Planets[i]
		state.solarSystem.Planets[i].Draw(renderer, 0, 255, 0, sdl.ALPHA_OPAQUE)

		pos := planet.Particle.Position
		vel := planet.Particle.Velocity
		vel.AddVec(pos)

		acc := planet.GravityTo(&state.solarSystem.Sun)
		acc.AddVec(pos)

		renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)

		arrow.NewArrow(
			sdl.Point{int32(pos.X), int32(pos.Y)},
			sdl.Point{int32(vel.X), int32(vel.Y)}, 15).Draw(renderer)

		renderer.SetDrawColor(255, 0, 0, sdl.ALPHA_OPAQUE)

		arrow.NewArrow(
			sdl.Point{int32(pos.X), int32(pos.Y)},
			sdl.Point{int32(acc.X), int32(acc.Y)}, 15).Draw(renderer)
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
		solarSystem: solsys.New(
			float64(winWidth)/2, float64(winHeight)/2,
			30, 100000)}

	state.solarSystem.SpawnPlanet(100, 30, 1, 10)

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
