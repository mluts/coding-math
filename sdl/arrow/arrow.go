package arrow

import (
	"../line"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Arrow struct {
	From, To sdl.Point
	width    float64
}

const angle float64 = math.Pi / 6

func NewArrow(From, To sdl.Point, width float64) Arrow {
	return Arrow{From, To, width}
}

func (a Arrow) Lines() [3]line.Line {
	return [3]line.Line{
		line.NewLine(a.From, a.To),
		a.leg().RotateFrom(angle),
		a.leg().RotateFrom(-angle)}
}

func (a Arrow) Rotate(origin sdl.Point, angle float64) Arrow {
	l := line.NewLine(a.From, a.To).RotateAround(origin, angle)
	return NewArrow(l.From, l.To, a.width)
}

func (a Arrow) Length() float64 {
	return line.NewLine(a.From, a.To).Length()
}

func (a Arrow) leg() line.Line {
	return line.NewLine(a.To, a.From).
		ScaleToLength(a.legLength())
}

func (a Arrow) legLength() float64 {
	return (a.width / 2) / math.Cos(math.Pi/6)
}

func (a Arrow) Draw(r *sdl.Renderer) {
	lines := a.Lines()
	for i := range a.Lines() {
		r.DrawLines([]sdl.Point{lines[i].From, lines[i].To})
	}
}
