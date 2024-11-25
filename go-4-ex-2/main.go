package main

import (
	"fmt"
	"math"
)

// computeHypotenuse calculates the hypotenuse of a triangle.
func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

type ShortSides struct {
	a, b float64
}

// Hypotenuse calculates the hypotenuse based on the struct.
func (s ShortSides) Hypotenuse() float64 {
	return computeHypotenuse(s.a, s.b)
}

func main() {
	fmt.Println(computeHypotenuse(3, 4))           // 5.000
	fmt.Println(ShortSides{6, 8}.Hypotenuse())     // 10.000
	fmt.Println(ShortSides{1.5, 2.0}.Hypotenuse()) // ~2.500
}
