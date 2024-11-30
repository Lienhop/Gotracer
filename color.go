package main

import "fmt"

type Color struct {
	r float64
	g float64
	b float64
}

func (c Color) writeColor() string {
	intensity := interval{0.000, 0.999}
	r := int(intensity.clamp(c.r) * 256)
	g := int(intensity.clamp(c.g) * 256)
	b := int(intensity.clamp(c.b) * 256)
	return fmt.Sprintf("%d %d %d\n", r, g, b)
}

func (c Color) add(w Color) Color {
	return Color{c.r + w.r, c.g + w.g, c.b + w.b}
}

func (c Color) scale(s float64) Color {
	return Color{c.r * s, c.g * s, c.b * s}
}
