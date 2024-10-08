package main

import (
	"fmt"
)

func main() {
	var r float64
	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	volume_125 := (4.0 / 3.0) * 3.1415926535 * r * r * r
	luas_125 := 4 * 3.1415926535 * r * r

	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume_125, luas_125)
}