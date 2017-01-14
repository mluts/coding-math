package solsys

import (
	"../particle"
	"../vector"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
	"math"
	"math/rand"
)

type Object struct {
	Particle part.Particle
	R, M     float64
}

type System struct {
	Sun     Object
	Planets []Object
}

func New(sunX float64, sunY float64, sunR float64, sunM float64) System {
	sun := Object{part.NewParticle(sunX, sunY, 0, 0, 0, 0), sunR, sunM}

	return System{Sun: sun}
}

func (s *System) SpawnPlanet(distanceToSun float64, radius float64, mass float64, speed float64) {
	angle := rand.Float64() * math.Pi * 2
	part := part.NewParticle(
		s.Sun.Particle.Position.X+distanceToSun,
		s.Sun.Particle.Position.Y, speed, 0, 0, 0)

	planet := Object{part, radius, mass}
	planet.Particle.Position.RotateAround(s.Sun.Particle.Position, angle)

	velAngle := planet.Particle.Position.AngleTo(s.Sun.Particle.Position) - math.Pi/2
	planet.Particle.Velocity.SetAngle(velAngle)

	s.Planets = append(s.Planets, planet)
}

func (s *System) Update(n float64) {
	for i := range s.Planets {
		s.Planets[i].update(&s.Sun, n)
	}
}

func (o *Object) update(sun *Object, n float64) {
	o.Particle.Acceleration = o.GravityTo(sun)
	o.Particle.Update(n)
}

func (o *Object) GravityTo(to *Object) vec.Vec2 {
	acc := vec.NewVec2(0, 0)
	d := o.Particle.Position.DistanceTo(to.Particle.Position)
	acc.SetLen(to.M / (d * d))
	acc.SetAngle(math.Pi + o.Particle.Position.AngleTo(to.Particle.Position))
	return acc
}

func (o *Object) Draw(renderer *sdl.Renderer, r, g, b, a uint8) {
	pos := o.Particle.Position
	gfx.FilledCircleRGBA(
		renderer, int(pos.X), int(pos.Y), int(o.R),
		r, g, b, a)
}
