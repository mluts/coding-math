package part

import (
	"../vector"
)

type Particle struct {
	Position, Velocity vec.Vec2
}

func NewParticle(X, Y, speed, angle float64) Particle {
	part := Particle{}
	part.Position = vec.NewVec2(X, Y)
	part.Velocity = vec.NewVec2(0, 0)
	part.Velocity.SetLen(speed)
	part.Velocity.SetAngle(angle)
	return part
}

func (p *Particle) Update(n float64) {
	vel := p.Velocity
	vel.Mul(n)
	p.Position.AddVec(vel)
}
