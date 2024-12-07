package main

import (
	"fmt"
	"math"
)

type Color struct {
	r float64
	g float64
	b float64
}

func linearToGamma(linearComp float64) float64 {
	if linearComp > 0 {
		return math.Sqrt(linearComp)
	}
	return 0
}

func (c Color) writeColor() string {
	intensity := interval{0.000, 0.999}

	r := linearToGamma(c.r)
	g := linearToGamma(c.g)
	b := linearToGamma(c.b)

	rByte := int(intensity.clamp(r) * 256)
	gByte := int(intensity.clamp(g) * 256)
	bByte := int(intensity.clamp(b) * 256)
	return fmt.Sprintf("%d %d %d\n", rByte, gByte, bByte)
}

func (c Color) add(w Color) Color {
	return Color{c.r + w.r, c.g + w.g, c.b + w.b}
}

func (c Color) scale(s float64) Color {
	return Color{c.r * s, c.g * s, c.b * s}
}
