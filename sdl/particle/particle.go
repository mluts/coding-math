package part

import (
	"../vector"
)

type Particle struct {
	Position, Velocity, Acceleration vec.Vec2
}

func NewParticle(X, Y, speed, angle, acc, accAngle float64) Particle {
	part := Particle{}
	part.Position = vec.NewVec2(X, Y)

	part.Velocity = vec.NewVec2(0, 0)
	part.Velocity.SetLen(speed)
	part.Velocity.SetAngle(angle)

	part.Acceleration = vec.NewVec2(0, 0)
	part.Acceleration.SetLen(acc)
	part.Acceleration.SetAngle(accAngle)
	return part
}

func (p *Particle) Update(n float64) {
	vel := p.Velocity
	vel.Mul(n)
	p.Position.AddVec(vel)

	acc := p.Acceleration
	acc.Mul(n)
	p.Velocity.AddVec(acc)
}
