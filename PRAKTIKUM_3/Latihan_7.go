package main

import (
	"fmt"
)

func main() {
    
	var bilangan int
	fmt.Print("Bilangan: ") // Meminta pengguna untuk memasukkan bilangan
	fmt.Scan(&bilangan) // Membaca input bilangan dari pengguna

	// Mengecek apakah bilangan lebih kecil atau sama dengan 1
	if bilangan <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1.")
		return // Menghentikan program jika bilangan tidak valid
	}

	// Mendapatkan semua faktor dari bilangan
	faktor := cariFaktor(bilangan)
	fmt.Print("Faktor: ")
	for _, f := range faktor {
		fmt.Printf("%d ", f) // Menampilkan setiap faktor
	}
	fmt.Println()

	// Mengecek apakah bilangan adalah bilangan prima
	if cariPrima(bilangan) {
		fmt.Println("Prima: true") // Menampilkan bahwa bilangan adalah prima
	} else {
		fmt.Println("Prima: false") // Menampilkan bahwa bilangan bukan prima
	}
}

// Fungsi untuk mencari faktor dari bilangan n
func cariFaktor(n int) []int {
	var faktor []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			faktor = append(faktor, i) // Menambahkan faktor ke dalam slice faktor
		}
	}
	return faktor // Mengembalikan slice yang berisi semua faktor
}

// Fungsi untuk mengecek apakah bilangan n adalah bilangan prima
func cariPrima(n int) bool {
	if n <= 1 {
		return false // Mengembalikan false jika bilangan kurang dari atau sama dengan 1
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false // Mengembalikan false jika ditemukan pembagi selain 1 dan bilangan itu sendiri
		}
	}
	return true // Mengembalikan true jika tidak ada pembagi selain 1 dan bilangan itu sendiri
}
