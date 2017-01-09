package tri

import (
	"../vector"
	// "fmt"
)

type Triangle struct {
	Center               vec.Vec2
	Height, Width, Angle float64
}

func NewTriangle(center vec.Vec2, h, w, a float64) Triangle {
	return Triangle{center, h, w, a}
}

func (t *Triangle) Points() [3]vec.Vec2 {
	top := vec.NewVec2(t.Center.X, t.Center.Y+t.Height/2)
	left := vec.NewVec2(t.Center.X-t.Width/2, t.Center.Y-t.Height/2)
	right := vec.NewVec2(t.Center.X+t.Width/2, t.Center.Y-t.Height/2)

	top.RotateAround(t.Center, t.Angle)
	left.RotateAround(t.Center, t.Angle)
	right.RotateAround(t.Center, t.Angle)

	return [3]vec.Vec2{top, left, right}
}
