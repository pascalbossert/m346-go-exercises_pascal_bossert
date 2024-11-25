package main

import (
	"errors"
	"fmt"
)

// computeGrade calculates the grade based on the achieved points and maximum points.
func computeGrade(gotPoints, maxPoints float64) (float64, error) {
	if gotPoints < 0 || maxPoints <= 0 || gotPoints > maxPoints {
		return 0, errors.New("invalid input values")
	}
	return 1.0 + 5.0*(gotPoints/maxPoints), nil
}

func main() {
	grades := []struct {
		got  float64
		max  float64
		desc string
	}{
		{17.5, 28.0, "Valid grade calculation"},
		{-5.0, 28.0, "Negative points error"},
		{30.0, 28.0, "Exceeded max points error"},
	}

	for _, grade := range grades {
		result, err := computeGrade(grade.got, grade.max)
		if err != nil {
			fmt.Printf("%s: Error: %v\n", grade.desc, err)
		} else {
			fmt.Printf("%s: Grade: %.3f\n", grade.desc, result)
		}
	}
}
