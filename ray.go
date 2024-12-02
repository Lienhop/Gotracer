package main

import "math"

var infinity = math.Inf(1)

type Ray struct {
	origin    Vec3
	direction Vec3
}

func (r Ray) at(t float64) Vec3 {
	return r.origin.add(r.direction.scale(t))
}

func rayColor(r Ray, world hittable, depth int) Color {

	if depth <= 0 {
		return Color{r: 0, g: 0, b: 0}
	}

	hitRecord := hitRecord{}
	if world.hit(r, interval{0.001, infinity}, &hitRecord) {
		direction := hitRecord.normal.add(randomUnitVector())
		return rayColor(Ray{hitRecord.p, direction}, world, depth-1).scale(0.5)
	}

	unitDirection := r.direction.unitVector()
	a := 0.5 * (unitDirection.y + 1.0)
	return Color{r: 1.0, g: 1.0, b: 1.0}.scale(1.0 - a).add(Color{r: 0.5, g: 0.7, b: 1.0}.scale(a))
}
