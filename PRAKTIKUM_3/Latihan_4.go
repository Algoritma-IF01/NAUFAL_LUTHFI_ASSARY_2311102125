package main

import (
	"fmt"
	"math"
)

// Fungsi Sqrt2 menghitung nilai aproksimasi dari akar 2 menggunakan rumus yang diberikan dengan parameter K sebagai batas iterasi
func Sqrt2(K int) float64 {
	// Inisialisasi hasil dengan nilai 1.0
	result := 1.0
	// Melakukan iterasi dari 0 hingga K untuk menghitung nilai hasil
	for k := 0; k <= K; k++ {
		// Menggunakan rumus (4k + 2)^2 / ((4k + 1) * (4k + 3))
		// Menghitung setiap suku dengan rumus (4k + 2)^2 / ((4k + 1) * (4k + 3))
		term := math.Pow(float64(4*k+2), 2) / (float64(4*k+1) * float64(4*k+3))
		// Mengalikan hasil dengan nilai suku yang dihitung
		result *= term
	}
	// Mengembalikan hasil akhir
	return result
}

// Fungsi main adalah titik masuk program, menerima input dari pengguna dan menghitung nilai aproksimasi akar 2
func main() {
	for i := 1; i <= 3; i++ {
	var K int
	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	// Menghitung nilai hampiran sqrt(2) menggunakan rumus yang diberikan
	result := Sqrt2(K)
	fmt.Printf("Nilai akar 2 = %.10f\n", result)
	}
}