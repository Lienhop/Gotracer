package main

import (
	"math"
	"math/rand"
)

func randomDouble() float64 {
	return rand.Float64()
}

func randomDoubleMinMax(min, max float64) float64 {
	return min + (max-min)*randomDouble()
}

func sampleSquare() Vec3 {
	return Vec3{randomDouble() - 0.5, randomDouble() - 0.5, 0}
}

func randomVec(min, max float64) Vec3 {
	return Vec3{randomDoubleMinMax(min, max), randomDoubleMinMax(min, max), randomDoubleMinMax(min, max)}
}

func randomUnitVector() Vec3 {
	for {
		p := randomVec(-1, 1)
		//1e160 check for floating point overflow
		if p.lengthSquared() >= 1 && 1e160 > p.lengthSquared() {
			continue
		}
		return p.divide(math.Sqrt(p.lengthSquared()))
	}
}

func (normal Vec3) randomOnHemisphere() Vec3 {
	onUnitSphere := randomUnitVector()
	if normal.dot(onUnitSphere) > 0.0 { // Same hemisphere as normal
		return onUnitSphere
	}
	return onUnitSphere.scale(-1)
}
