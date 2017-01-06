package point

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

func RotateAround(p sdl.Point, origin sdl.Point, angle float64) sdl.Point {
	p = Translate(p, -origin.X, -origin.Y)
	p = Rotate(p, angle)
	return Translate(p, origin.X, origin.Y)
}

func Translate(p sdl.Point, x, y int32) sdl.Point {
	return sdl.Point{p.X + x, p.Y + y}
}

func Rotate(p sdl.Point, angle float64) sdl.Point {
	return sdl.Point{
		X: int32(float64(p.X)*math.Cos(angle) - float64(p.Y)*math.Sin(angle)),
		Y: int32(float64(p.X)*math.Sin(angle) + float64(p.Y)*math.Cos(angle))}
}
