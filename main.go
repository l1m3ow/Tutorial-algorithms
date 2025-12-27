package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Acos(x) - math.Sqrt(1-0.3*math.Pow(x, 3))
}

func Bisection(a, b, epsilon float64) float64 {
	if math.Abs(b-a) < epsilon {
		return (a + b) / 2
	}

	c := (a + b) / 2
	fc := f(c)

	if fc == 0 {
		return c
	}

	if f(a)*fc < 0 {
		return Bisection(a, c, epsilon)
	} else {
		return Bisection(c, b, epsilon)
	}
}

func main() {
	a := 0.0
	b := 1.0
	epsilon := 0.000001
	fmt.Printf("На отрезке [0;1] корень заданной функции = %f", Bisection(a, b, epsilon))
}
