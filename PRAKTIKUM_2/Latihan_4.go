package main

import (
	"fmt"
)

func main() {
	var celsius_125 float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius_125)

	reamur_125 := celsius_125 * 4 / 5
	fahrenheit_125 := (celsius_125 * 9 / 5) + 32
	kelvin_125 := (fahrenheit_125 + 459.67) * 5 / 9

	fmt.Printf("Derajat Reamur: %.0f\n", reamur_125)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit_125)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin_125)
}