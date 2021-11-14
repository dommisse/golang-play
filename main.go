package main

import (
	"fmt"
	"math"
)

func main() {
	op := math.Abs
	a := []float64{1, -2}
	b := Map(op, a)
	fmt.Println(b) // [1 2]

	c := Map(func(x float64) float64 { return 10 * x }, b)
	fmt.Println(c) // [10, 20]
}
