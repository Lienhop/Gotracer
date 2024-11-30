package main

type interval struct {
	min, max float64
}

func (i interval) size() float64 {
	return i.max - i.min
}

func (i interval) contains(x float64) bool {
	return x >= i.min && x <= i.max
}

func (i interval) surrounds(x float64) bool {
	return x > i.min && x < i.max
}

func (i interval) clamp(x float64) float64 {
	if x < i.min {
		return i.min
	}
	if x > i.max {
		return i.max
	}
	return x
}
