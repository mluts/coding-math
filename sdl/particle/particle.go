package part

import (
	"../vector"
	"math"
)

type Particle struct {
	Position, Velocity, Acceleration vec.Vec2
	Friction                         float64
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

	part.Friction = 0
	return part
}

func (p *Particle) Update(n float64) {
	p.doFriction(n)
	p.move(n)
	p.accelerate(n)
}

func (p *Particle) HitEdgesAsCircle(r, edgeX, edgeY float64) {
	pos := &p.Position
	vel := &p.Velocity
	rate := 0.8

	if (pos.X - r) < 0 {
		pos.X = r
		vel.X *= -rate
	} else if pos.X+r > edgeX {
		pos.X = edgeX - r
		vel.X *= -rate
	}

	if (pos.Y - r) < 0 {
		pos.Y = r
		vel.Y *= -rate
	} else if pos.Y+r > edgeY {
		pos.Y = edgeY - r
		vel.Y *= -rate
	}
}

func (p *Particle) doFriction(n float64) {
	friction := vec.NewVec2(0, 0)
	friction.SetLen(p.Friction)
	friction.SetAngle(math.Pi + p.Velocity.Angle())
	friction.Mul(n)

	if friction.Len() >= p.Velocity.Len() {
		p.Velocity.SetLen(0)
	} else {
		p.Velocity.AddVec(friction)
	}
}

func (p *Particle) move(n float64) {
	vel := p.Velocity
	vel.Mul(n)
	p.Position.AddVec(vel)
}

func (p *Particle) accelerate(n float64) {
	acc := p.Acceleration
	acc.Mul(n)
	p.Velocity.AddVec(acc)
}
