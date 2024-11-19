package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	// Array untuk menampung data berat ikan dengan kapasitas maksimum 1000
	if x > 1000 {
		fmt.Println("Jumlah ikan tidak boleh lebih dari 1000")
		return
	}
	
	beratIkan := make([]float64, x)

	// Membaca berat setiap ikan
	fmt.Println("Masukkan berat ikan satu per satu:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan %d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	// Menghitung jumlah wadah
	jumlahWadah := (x + y - 1) / y // Pembulatan ke atas jika tidak terbagi rata

	// Menginisialisasi array wadah untuk menyimpan berat setiap wadah
	wadah := make([]float64, jumlahWadah)

	// Memasukkan ikan ke dalam wadah secara berurutan
	for i := 0; i < x; i++ {
		idxWadah := i / y
		wadah[idxWadah] += beratIkan[i]
	}

	// Menampilkan total berat ikan di setiap wadah
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %.2f\n", i+1, wadah[i])
	}

	// Menghitung rata-rata berat per wadah
	totalBerat := 0.0
	for i := 0; i < jumlahWadah; i++ {
		totalBerat += wadah[i]
	}
	rataRata := totalBerat / float64(jumlahWadah)

	// Menampilkan rata-rata berat ikan per wadah
	fmt.Printf("\nRata-rata berat ikan di setiap wadah: %.2f\n", rataRata)
}
