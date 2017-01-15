package main

import (
	"../sdl/particle"
	// "fmt"
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
	running      bool
	lastUpdateMs uint32
	speedMs      int
	particles    []part.Particle
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
	currentTick := sdl.GetTicks()
	dt := float64(currentTick-state.lastUpdateMs) / 1000
	speed := int(1)
	precision := int(10)

	for k := 0; k < speed; k++ {
		for i := range state.particles {
			p := &state.particles[i]
			for j := 0; j < precision; j++ {
				p.Update(dt / float64(precision))
				p.HitEdgesAsCircle(5, float64(winWidth), float64(winHeight))
			}
		}
	}
	// fmt.Printf("p0 vel %f\n", state.particles[0].Velocity.Len())
}

func draw(state *State, renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)

	for _, p := range state.particles {
		gfx.FilledCircleRGBA(
			renderer,
			int(p.Position.X),
			int(p.Position.Y),
			5,
			0, 0, 0, sdl.ALPHA_OPAQUE)
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
		lastUpdateMs: 0}

	for i := 0; i < 50; i++ {
		p := part.NewParticle(
			float64(winWidth/2)+float64(rand.Intn(50)-25),
			float64(winHeight*3/10)+float64(rand.Intn(50)-25),
			rand.Float64()*500, rand.Float64()*math.Pi*2,
			150, math.Pi/2)
		p.Friction = 25

		state.particles = append(state.particles, p)
	}

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
