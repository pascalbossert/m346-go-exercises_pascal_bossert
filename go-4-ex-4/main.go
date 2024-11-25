package main

import "fmt"

// Conversion functions
func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// Types for methods
type Celsius float64
type Fahrenheit float64

func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Println(convertCelsiusToFahrenheit(23))   // 73.4
	fmt.Println(convertFahrenheitToCelsius(73.4)) // ~23.0

	var tempC Celsius = 23
	var tempF Fahrenheit = tempC.ConvertToFahrenheit()
	fmt.Println(tempF)                    // 73.4
	fmt.Println(tempF.ConvertToCelsius()) // ~23.0
}
