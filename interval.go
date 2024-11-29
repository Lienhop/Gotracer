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
