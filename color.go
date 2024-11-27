package main

import "fmt"

type Color struct {
	r float64
	g float64
	b float64
}

func (c Color) writeColor() string {
	return fmt.Sprintf("%d %d %d\n", int(c.r*255), int(c.g*255), int(c.b*255))
}

func (c Color) add(w Color) Color {
	return Color{c.r + w.r, c.g + w.g, c.b + w.b}
}

func (c Color) scale(s float64) Color {
	return Color{c.r * s, c.g * s, c.b * s}
}
