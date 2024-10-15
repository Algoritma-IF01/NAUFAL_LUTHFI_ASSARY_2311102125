package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 4; i++ {
	var berat1, berat2 float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		// Cek jika salah satu kantong beratnya negatif
		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Cek jika total berat kedua kantong melebihi 150 kg
		if berat1+berat2 > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Cek apakah selisih berat di kedua kantong lebih dari atau sama dengan 9 kg
		if math.Abs(berat1-berat2) >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
		}
	}
}