package main

import (
	"fmt"
	"math"
)

// computeDiscriminant calculates the discriminant of the quadratic equation.
func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

// computeQuadraticFormula solves the quadratic equation using the discriminant.
func computeQuadraticFormula(a, b, c float64) []float64 {
	d := computeDiscriminant(a, b, c)
	if d > 0 {
		root1 := (-b + math.Sqrt(d)) / (2 * a)
		root2 := (-b - math.Sqrt(d)) / (2 * a)
		return []float64{root1, root2}
	} else if d == 0 {
		return []float64{-b / (2 * a)}
	}
	return nil
}

func main() {
	fmt.Println(computeQuadraticFormula(3, 4, 1)) // Two solutions
	fmt.Println(computeQuadraticFormula(2, 4, 2)) // One solution
	fmt.Println(computeQuadraticFormula(3, 4, 2)) // No solution
}
