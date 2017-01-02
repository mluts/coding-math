package main

import (
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
	running                   bool
	lastUpdateMs              uint32
	speedMs                   int
	lastTick                  uint32
	angle, speed              float64
	circleX, circleY, circleR int
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
		case sdl.K_PLUS:
			state.speedMs += 100
		case sdl.K_MINUS:
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
	state.angle += state.speed
}

func draw(state *State, renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	gfx.FilledCircleRGBA(
		renderer, state.circleX, state.circleY+int(math.Sin(state.angle)*float64(winHeight*2/5)),
		state.circleR, 0, 0, 0, sdl.ALPHA_OPAQUE)
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
		speedMs:      10,
		lastUpdateMs: 0,
		angle:        0,
		circleX:      winWidth / 2,
		circleY:      winHeight / 2,
		circleR:      80,
		speed:        0.1}

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
