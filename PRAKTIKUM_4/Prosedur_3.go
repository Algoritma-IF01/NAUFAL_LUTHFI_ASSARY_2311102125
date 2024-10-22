package main

import (
	"fmt"
)

// Procedure cetakDeret mencetak deret bilangan sesuai aturan Skiena
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Printf("%d\n", n) // Cetak 1 sebagai elemen terakhir deret
}

func main() {
	var n int
	fmt.Print("Masukkan nilai awal): ")
	fmt.Scan(&n)
	if n > 0 && n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Println("Nilai harus berupa bilangan positif")
	}
}
