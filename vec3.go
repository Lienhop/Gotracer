package main

import "math"

type Vec3 struct {
	x, y, z float64
}

func (v Vec3) toColor() Color {
	return Color{v.x, v.y, v.z}
}

func (v Vec3) subtract(w Vec3) Vec3 {
	return Vec3{v.x - w.x, v.y - w.y, v.z - w.z}
}

func (v Vec3) add(w Vec3) Vec3 {
	return Vec3{v.x + w.x, v.y + w.y, v.z + w.z}
}

func (v Vec3) scale(s float64) Vec3 {
	return Vec3{v.x * s, v.y * s, v.z * s}
}

func (v Vec3) divide(d float64) Vec3 {
	return Vec3{v.x / d, v.y / d, v.z / d}
}

func (v Vec3) lengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) length() float64 {
	return math.Sqrt(v.lengthSquared())
}

func (v Vec3) dot(w Vec3) float64 {
	return v.x*w.x + v.y*w.y + v.z*w.z
}

func (v Vec3) cross(w Vec3) Vec3 {
	return Vec3{v.y*w.z - v.z*w.y, v.z*w.x - v.x*w.z, v.x*w.y - v.y*w.x}
}

func (v Vec3) unitVector() Vec3 {
	return v.divide(v.length())
}

func (v Vec3) nearZero() bool {
	s := 1e-8
	return math.Abs(v.x) < s && math.Abs(v.y) < s && math.Abs(v.z) < s
}

func (v Vec3) reflect(n Vec3) Vec3 {
	return v.subtract(n.scale(2 * v.dot(n)))
}
