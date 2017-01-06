package sdlext

import (
	"./arrow"
	"./line"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderer sdl.Renderer

func (r *Renderer) DrawLine(l line.Line) {
	ps := []sdl.Point{l.From, l.To}
	(*sdl.Renderer)(r).DrawLines(ps)
}

func (r Renderer) DrawArrow(a arrow.Arrow) {
	for _, l := range a.Lines() {
		r.DrawLine(l)
	}
}
