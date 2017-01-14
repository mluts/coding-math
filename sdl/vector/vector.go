package vec

import (
	"math"
)

type Vec2 struct {
	X, Y float64
}

func NewVec2(X, Y float64) Vec2 {
	return Vec2{X, Y}
}

func (v *Vec2) AddVec(b Vec2) {
	v.X += b.X
	v.Y += b.Y
}

func (v *Vec2) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) SetLen(len float64) {
	angle := v.Angle()
	v.X = math.Cos(angle) * len
	v.Y = math.Sin(angle) * len
}

func (v *Vec2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v *Vec2) SetAngle(angle float64) {
	len := v.Len()
	v.X = math.Cos(angle) * len
	v.Y = math.Sin(angle) * len
}

func (v *Vec2) Mul(n float64) {
	v.X *= n
	v.Y *= n
}

func (v *Vec2) Div(n float64) {
	v.X /= n
	v.Y /= n
}

func (v *Vec2) Translate(x, y float64) {
	v.X += x
	v.Y += y
}

func (v *Vec2) RotateAround(around Vec2, angle float64) {
	v.Translate(-around.X, -around.Y)
	x, y := v.X, v.Y
	v.X = x*math.Cos(angle) - y*math.Sin(angle)
	v.Y = x*math.Sin(angle) + y*math.Cos(angle)
	v.Translate(around.X, around.Y)
}

func (v *Vec2) AngleTo(to Vec2) float64 {
	v.Translate(-to.X, -to.Y)
	angle := v.Angle()
	v.Translate(to.X, to.Y)
	return angle
}

func (v *Vec2) DistanceTo(to Vec2) float64 {
	x := v.X - to.X
	y := v.Y - to.Y
	return math.Sqrt(x*x + y*y)
}
