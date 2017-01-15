package bez

import (
	"../vector"
	"math"
)

func LinearBezier(p0, p1 vec.Vec2, t float64) vec.Vec2 {
	x := p0.X + (p1.X-p0.X)*t
	y := p0.Y + (p1.Y-p0.Y)*t
	return vec.NewVec2(x, y)
}

func QuadraticBezier(p0, p1, p2 vec.Vec2, t float64) vec.Vec2 {
	x := math.Pow(1-t, 2)*p0.X +
		(1-t)*2*t*p1.X +
		math.Pow(t, 2)*p2.X

	y := math.Pow(1-t, 2)*p0.Y +
		(1-t)*2*t*p1.Y +
		math.Pow(t, 2)*p2.Y

	return vec.NewVec2(x, y)
}
