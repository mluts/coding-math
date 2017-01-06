package line

import (
	"../point"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Line struct {
	From, To sdl.Point
}

func NewLine(From, To sdl.Point) Line {
	return Line{From, To}
}

func (l Line) Center() sdl.Point {
	return sdl.Point{
		X: (l.From.X + l.To.X) / 2,
		Y: (l.From.Y + l.To.Y) / 2}
}

func (l Line) RotateFrom(angle float64) Line {
	return l.RotateAround(l.From, angle)
}

func (l Line) RotateTo(angle float64) Line {
	return l.RotateAround(l.To, angle)
}

func (l Line) RotateAround(p sdl.Point, angle float64) Line {
	return Line{
		point.RotateAround(l.From, p, angle),
		point.RotateAround(l.To, p, angle)}
}

func (l Line) Translate(x, y int32) Line {
	return Line{
		point.Translate(l.From, x, y),
		point.Translate(l.To, x, y)}
}

func (l Line) Length() float64 {
	x := float64(l.To.X - l.From.X)
	y := float64(l.To.Y - l.From.Y)

	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func (l Line) ScaleToLength(len float64) Line {
	lineLength := l.Length()

	To := point.Translate(l.To, -l.From.X, -l.From.Y)

	To.X = int32(float64(To.X) * len / lineLength)
	To.Y = int32(float64(To.Y) * len / lineLength)

	To = point.Translate(To, l.From.X, l.From.Y)

	return NewLine(l.From, To)
}
