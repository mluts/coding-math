package main

import (
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
	lastTick     uint32
	flying       []FlyingObject
}

type Circle struct {
	c      sdl.Point
	radius int
	color  sdl.Color
	filled bool
}

type Drawable interface {
	getCenter() sdl.Point
	setCenter(sdl.Point) Drawable
	draw(*sdl.Renderer)
}

type FlyingObject struct {
	obj                            Drawable
	xangle, yangle, xspeed, yspeed float64
}

func newFlyingObject(d Drawable, xa, ya, xs, ys float64) FlyingObject {
	return FlyingObject{d, xa, ya, xs, ys}
}

func (f FlyingObject) draw(r *sdl.Renderer) {
	c := f.obj.getCenter()

	o := f.obj.setCenter(sdl.Point{
		c.X + int32(math.Cos(f.xangle)*float64(200)),
		c.Y + int32(math.Sin(f.yangle)*float64(150))})

	o.draw(r)
}

func (f *FlyingObject) fly() {
	f.xangle += f.xspeed
	f.yangle += f.yspeed
}

func newCircle(center sdl.Point, radius int, color sdl.Color, filled bool) Circle {
	return Circle{center, radius, color, filled}
}

func (c Circle) draw(r *sdl.Renderer) {
	if c.filled {
		gfx.FilledCircleColor(r, int(c.c.X), int(c.c.Y), c.radius, c.color)
	} else {
		gfx.CircleColor(r, int(c.c.X), int(c.c.Y), c.radius, c.color)
	}
}

func (c Circle) move(p sdl.Point) Circle {
	return Circle{
		p, c.radius, c.color, c.filled}
}

func (c Circle) getCenter() sdl.Point {
	return c.c
}

func (c Circle) setCenter(p sdl.Point) Drawable {
	return c.move(p)
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
	}
}

func updateState(state *State) {
	for i, o := range state.flying {
		o.fly()
		state.flying[i] = o
	}
}

func draw(state *State, renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)
	for _, o := range state.flying {
		o.draw(renderer)
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
		speedMs:      100,
		lastUpdateMs: 0,
		flying:       make([]FlyingObject, 0)}

	for i := 0; i < 1000; i++ {
		color := sdl.Color{0, 0, 0, sdl.ALPHA_OPAQUE}
		throw := int((float64(winWidth) * 0.7) / 2)
		c := newCircle(
			sdl.Point{
				int32(winWidth/2 + rand.Intn(throw) - throw/2),
				int32(winHeight/2 + rand.Intn(throw) - throw/2)},
			int(1+rand.Float64()*5),
			color,
			true)
		state.flying = append(state.flying, newFlyingObject(
			c, float64(rand.Intn(50)-25), float64(rand.Intn(50)-25),
			0.01+rand.Float64()*0.05, 0.01+rand.Float64()*0.05))
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
