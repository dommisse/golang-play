package main

import (
	"dj.com/functions"
	_ "dj.com/functions"
	"fmt"
	"math"
)

func main() {
	op := math.Abs
	a := []float64{1, -2}

	b := functions.Map(op, a)
	fmt.Println(b) // [1 2]

	c := functions.Map(func(x float64) float64 { return 10 * x }, b)
	fmt.Println(c) // [10, 20]
}
