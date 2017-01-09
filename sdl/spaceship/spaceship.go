package sp

import (
	"../triangle"
	"../vector"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
)

type Spaceship struct {
	Tri                    tri.Triangle
	Velocity, Acceleration vec.Vec2
}

func NewSpaceship(pos vec.Vec2, h, w float64) Spaceship {
	tri := tri.NewTriangle(pos, h, w, 0)
	return Spaceship{tri, vec.NewVec2(0, 0), vec.NewVec2(0, 0)}
}

func (s *Spaceship) Draw(r *sdl.Renderer) {
	var points [3]vec.Vec2 = s.Tri.Points()

	gfx.TrigonRGBA(
		r,
		int(points[0].X),
		int(points[0].Y),
		int(points[1].X),
		int(points[1].Y),
		int(points[2].X),
		int(points[2].Y),
		0, 0, 0, sdl.ALPHA_OPAQUE)
}

func (s *Spaceship) DrawTail(r *sdl.Renderer) {
	top := s.Tri.Center
	bottom := s.Tri.Center
	bottom.Y -= s.Tri.Height
	top.RotateAround(s.Tri.Center, s.Tri.Angle)
	bottom.RotateAround(s.Tri.Center, s.Tri.Angle)

	r.DrawLines([]sdl.Point{
		sdl.Point{int32(top.X), int32(top.Y)},
		sdl.Point{int32(bottom.X), int32(bottom.Y)}})
}

func (s *Spaceship) GetAcceleration() vec.Vec2 {
	return s.Acceleration
}

func (s *Spaceship) SetAcceleration(acc vec.Vec2) {
	s.Acceleration = acc
}

func (s *Spaceship) GetVelocity() vec.Vec2 {
	return s.Velocity
}

func (s *Spaceship) SetVelocity(vel vec.Vec2) {
	s.Velocity = vel
}

func (s *Spaceship) GetPosition() vec.Vec2 {
	return s.Tri.Center
}

func (s *Spaceship) SetPosition(pos vec.Vec2) {
	s.Tri.Center = pos
}

func (s *Spaceship) Rotate(angle float64) {
	s.Tri.Angle += angle
}
