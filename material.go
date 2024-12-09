package main

import (
	"math"
)

type Material interface {
	scatter(rIn Ray, rec *hitRecord, attenuation *Color, scattered *Ray) bool
}

type Lambertian struct {
	albedo Color
}

func (l Lambertian) scatter(rIn Ray, rec *hitRecord, attenuation *Color, scattered *Ray) bool {
	scatterDirection := rec.normal.add(randomUnitVector())
	if scatterDirection.nearZero() {
		scatterDirection = rec.normal
	}
	*scattered = Ray{rec.p, scatterDirection}
	*attenuation = l.albedo
	return true
}

type Metal struct {
	albedo Color
	fuzz   float64
}

func (m Metal) scatter(rIn Ray, rec *hitRecord, attenuation *Color, scattered *Ray) bool {
	reflected := rIn.direction.unitVector().reflect(rec.normal)
	reflected = reflected.unitVector().add(randomUnitVector().scale(m.fuzz))
	*scattered = Ray{rec.p, reflected}
	*attenuation = m.albedo
	return rec.normal.dot(scattered.direction) > 0
}

type Dielectric struct {
	refractionIndex float64
}

func (d Dielectric) scatter(rIn Ray, rec *hitRecord, attenuation *Color, scattered *Ray) bool {
	*attenuation = Color{r: 1.0, g: 1.0, b: 1.0}
	var ri float64

	if rec.frontFace {
		ri = 1.0 / d.refractionIndex
	} else {
		ri = d.refractionIndex
	}

	unitDirection := rIn.direction.unitVector()

	cosTheta := math.Min(rec.normal.dot(unitDirection.scale(-1)), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := ri*sinTheta > 1.0

	var direction Vec3

	if cannotRefract || reflectance(cosTheta, ri) > randomDouble() {
		direction = unitDirection.reflect(rec.normal)
	} else {
		direction = unitDirection.refract(rec.normal, ri)
	}

	*scattered = Ray{rec.p, direction}
	return true
}

// Polynomial Schlick approximation for reflectivity
func reflectance(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
