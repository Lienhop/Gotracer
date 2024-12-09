package main

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
