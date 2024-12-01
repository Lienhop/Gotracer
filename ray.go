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

func rayColor(r Ray, world hittable) Color {
	hitRecord := hitRecord{}
	if world.hit(r, interval{0, infinity}, &hitRecord) {
		direction := hitRecord.normal.randomOnHemisphere()
		return rayColor(Ray{hitRecord.p, direction}, world).scale(0.5)
		//return hitRecord.normal.add(Vec3{1, 1, 1}).scale(0.5).toColor()
	}

	unitDirection := r.direction.unitVector()
	a := 0.5 * (unitDirection.y + 1.0)
	return Color{r: 1.0, g: 1.0, b: 1.0}.scale(1.0 - a).add(Color{r: 0.5, g: 0.7, b: 1.0}.scale(a))
}
