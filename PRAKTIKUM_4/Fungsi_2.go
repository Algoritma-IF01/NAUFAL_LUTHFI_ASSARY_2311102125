package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Fungsi untuk menghitung permutasi
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	// Meminta input dari pengguna
	fmt.Println("Masukkan nilai a, b, c, d:")
	fmt.Scan(&a, &b, &c, &d)

	// Menghitung permutasi dan kombinasi untuk a terhadap c
	p1 := permutation(a, c)
	c1 := combination(a, c)

	// Menghitung permutasi dan kombinasi untuk b terhadap d
	p2 := permutation(b, d)
	c2 := combination(b, d)

	// Output hasil 
	fmt.Printf("P(%d,%d) = %d\n", a, c, p1)
	fmt.Printf("C(%d,%d) = %d\n", a, c, c1)
	fmt.Printf("P(%d,%d) = %d\n", b, d, p2)
	fmt.Printf("C(%d,%d) = %d\n", b, d, c2)
}
