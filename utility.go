package main

import (
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
