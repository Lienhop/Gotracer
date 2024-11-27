package main

type Ray struct {
	origin    Vec3
	direction Vec3
}

func (r Ray) at(t float64) Vec3 {
	return r.origin.add(r.direction.scale(t))
}
