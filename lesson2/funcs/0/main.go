package main

import (
	"fmt"
)

const epsilon = 0.01
const start = 2

// TODO: Реализовать вычисление Квадратного корня
func Sqrt(x float64) float64 {
	z := newton(start, x)
	for zn, delta := z, z; delta > epsilon; z = zn {
		zn = newton(z, x)
		delta = z - zn
	}
	return z
}

func newton(z, x float64) float64 {
	return z - (z*z-x)/(2*z)
}

func main() {
	fmt.Println(Sqrt(2))
}
