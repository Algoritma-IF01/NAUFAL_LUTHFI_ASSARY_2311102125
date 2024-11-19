package main

import (
	"fmt"
)

type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum balita
func hitungMinMax(arrBerat arrBalita, jumlahData int, bMin *float64, bMax *float64) {
	// Inisialisasi bMin dan bMax dengan nilai pertama dari array
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	// Loop untuk mencari nilai minimum dan maksimum
	for i := 1; i < jumlahData; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func rerata(arrBerat arrBalita, jumlahData int) float64 {
	var total float64 = 0.0

	// Loop untuk menjumlahkan semua berat
	for i := 0; i < jumlahData; i++ {
		total += arrBerat[i]
	}

	// Menghitung rata-rata
	return total / float64(jumlahData)
}

func main() {
	var jumlahBalita int
	var arrBerat arrBalita

	// Input jumlah data balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&jumlahBalita)

	// Memastikan jumlah data tidak lebih dari kapasitas array
	if jumlahBalita > 100 {
		fmt.Println("Jumlah data tidak boleh lebih dari 100")
		return
	}

	// Input berat masing-masing balita
	for i := 0; i < jumlahBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	// Deklarasi variabel untuk menyimpan nilai minimum dan maksimum
	var beratMin, beratMax float64

	// Memanggil fungsi hitungMinMax untuk mencari nilai minimum dan maksimum
	hitungMinMax(arrBerat, jumlahBalita, &beratMin, &beratMax)

	// Memanggil fungsi rerata untuk menghitung rata-rata
	rataRata := rerata(arrBerat, jumlahBalita)

	// Menampilkan hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}
