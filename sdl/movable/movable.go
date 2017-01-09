package movable

import "../vector"

type Movable2 interface {
	GetAcceleration() vec.Vec2
	GetVelocity() vec.Vec2
	GetPosition() vec.Vec2
	SetAcceleration(vec.Vec2)
	SetVelocity(vec.Vec2)
	SetPosition(vec.Vec2)
}

func Move2(m Movable2, n float64) {
	move(m, n)
	accelerate(m, n)
}

func move(m Movable2, n float64) {
	vel := m.GetVelocity()
	vel.Mul(n)
	pos := m.GetPosition()
	pos.AddVec(vel)
	m.SetPosition(pos)
}

func accelerate(m Movable2, n float64) {
	acc := m.GetAcceleration()
	acc.Mul(n)
	vel := m.GetVelocity()
	vel.AddVec(acc)
	m.SetVelocity(vel)
}
