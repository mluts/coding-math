package main

import (
	"../sdl/movable"
	"../sdl/spaceship"
	"../sdl/vector"
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
	running                           bool
	lastUpdateMs                      uint32
	speedMs                           int
	spaceship                         sp.Spaceship
	leftKey, rightKey, upKey, downKey bool
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
		case sdl.K_UP:
			state.upKey = true
		case sdl.K_DOWN:
			state.downKey = true
		case sdl.K_LEFT:
			state.leftKey = true
		case sdl.K_RIGHT:
			state.rightKey = true
		}

	case *sdl.KeyUpEvent:
		switch t.Keysym.Sym {
		case sdl.K_UP:
			state.upKey = false
		case sdl.K_DOWN:
			state.downKey = false
		case sdl.K_LEFT:
			state.leftKey = false
		case sdl.K_RIGHT:
			state.rightKey = false
		}
	}
}

func updateState(state *State) {
	var (
		currentTick         = sdl.GetTicks()
		dt          float64 = float64(currentTick-state.lastUpdateMs) / 1000
		mov                 = &state.spaceship
	)

	if state.leftKey {
		state.spaceship.Rotate(-dt * math.Pi)
	}

	if state.rightKey {
		state.spaceship.Rotate(dt * math.Pi)
	}

	if state.upKey {
		state.spaceship.Acceleration.SetAngle(state.spaceship.Tri.Angle + math.Pi/2)
		state.spaceship.Acceleration.SetLen(300)
	} else {
		state.spaceship.Acceleration.SetLen(0)
	}

	if state.downKey && !state.upKey {
		state.spaceship.Velocity = vec.NewVec2(0, 0)
	}

	movable.Move2(mov, dt)

	pos := mov.GetPosition()
	if pos.X < 0 {
		pos.X = float64(winWidth)
	} else if pos.X > float64(winWidth) {
		pos.X = 0
	}

	if pos.Y < 0 {
		pos.Y = float64(winHeight)
	} else if pos.Y > float64(winHeight) {
		pos.Y = 0
	}

	mov.SetPosition(pos)
}

func draw(state *State, renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)

	state.spaceship.Draw(renderer)
	if state.spaceship.Acceleration.Len() > 0 {
		state.spaceship.DrawTail(renderer)
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
		spaceship: sp.NewSpaceship(
			vec.NewVec2(float64(winWidth/2), float64(winHeight/2)),
			30, 20)}

	state.spaceship.Rotate(math.Pi)

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
