package main

import (
	"math"
)

type hitRecord struct {
	p         Vec3
	normal    Vec3
	t         float64
	frontFace bool
}

func (h *hitRecord) setFaceNormal(r Ray, outwardNormal Vec3) {
	h.frontFace = r.direction.dot(outwardNormal) < 0
	if h.frontFace {
		h.normal = outwardNormal
	} else {
		h.normal = outwardNormal.scale(-1)
	}
}

type hittable interface {
	hit(r Ray, tMin float64, tMax float64, rec *hitRecord) bool
}

type hittableList struct {
	list []hittable
}

func (l *hittableList) add(h hittable) {
	l.list = append(l.list, h)
}

func (l *hittableList) clear() {
	l.list = nil
}

func (l hittableList) hit(r Ray, tMin float64, tMax float64, rec *hitRecord) bool {
	tempRec := hitRecord{}
	hitAnything := false
	closestSoFar := tMax
	for _, object := range l.list {
		if object.hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			*rec = tempRec
		}
	}
	return hitAnything
}

type Sphere struct {
	center Vec3
	radius float64
}

func (s Sphere) hit(r Ray, tMin float64, tMax float64, rec *hitRecord) bool {
	oc := s.center.subtract(r.origin)
	a := r.direction.lengthSquared()
	h := r.direction.dot(oc)
	c := oc.lengthSquared() - s.radius*s.radius
	discriminant := h*h - a*c

	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	root := (h - sqrtd) / a
	if root <= tMin || root >= tMax {
		root = (h + sqrtd) / a
		if root <= tMin || root >= tMax {
			return false
		}
	}
	rec.t = root
	rec.p = r.at(rec.t)
	outwardNormal := rec.p.subtract(s.center).divide(s.radius)
	rec.setFaceNormal(r, outwardNormal)

	return true
}
