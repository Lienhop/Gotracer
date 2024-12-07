package main

type Material interface {
	scatter(rIn Ray, rec *hitRecord, attenuation *Color, scattered *Ray) bool
}

type Lambertian struct {
	albedo Color
}

func (l Lambertian) scatter(rIn Ray, rec hitRecord, attenuation *Color, scattered *Ray) bool {
	scatterDirection := rec.normal.add(randomUnitVector())
	if scatterDirection.nearZero() {
		scatterDirection = rec.normal
	}
	*scattered = Ray{rec.p, scatterDirection}
	*attenuation = l.albedo
	return true
}
