package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan elemen array menggunakan algoritma selection sort
func selectionSort(arr []int) []int {
    n := len(arr) // Menyimpan panjang array
    for i := 0; i < n-1; i++ { // Iterasi elemen array
        minIdx := i // Anggap elemen pertama sebagai elemen terkecil
        for j := i + 1; j < n; j++ { // Bandingkan dengan elemen lainnya
            if arr[j] < arr[minIdx] { // Jika ada elemen yang lebih kecil
                minIdx = j // Perbarui indeks elemen terkecil
            }
        }
        // Tukar posisi antara elemen terkecil dan elemen di posisi i
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
    return arr // Mengembalikan array yang sudah terurut
}

func main() {
	var n int // Variabel untuk jumlah daerah
	fmt.Println("Masukkan jumlah daerah kerabat:")
	fmt.Scan(&n) // Input jumlah daerah

	// Validasi jumlah daerah
	if n <= 0 || n >= 1000 {
		fmt.Println("Jumlah daerah harus di antara 1 dan 999.")
		return
	}

	// Array untuk menyimpan nomor rumah dari setiap daerah
	var daerahRumah [][]int

	// Proses input data rumah dari tiap daerah
	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah dan nomor rumah untuk daerah %d:\n", i+1)
		fmt.Scan(&m) // Input jumlah rumah di daerah ini

		// Validasi jumlah rumah
		if m <= 0 || m >= 1000000 {
			fmt.Println("Jumlah rumah kerabat harus di antara 1 dan 999999.")
			return
		}

		rumah := make([]int, m) // Membuat array untuk nomor rumah
		for j := 0; j < m; j++ { // Input nomor rumah
			fmt.Scan(&rumah[j])
		}

		// Menambahkan data rumah ke array daerahRumah
		daerahRumah = append(daerahRumah, rumah)
	}

	// Menampilkan hasil setelah input selesai
	fmt.Println("\nKeluaran:")
	for _, rumah := range daerahRumah {
		// Mengurutkan rumah menggunakan fungsi selectionSort
		sortedRumah := selectionSort(rumah)
		for _, num := range sortedRumah {
			fmt.Printf("%d ", num) // Menampilkan rumah yang sudah diurutkan
		}
		fmt.Println() // Pindah baris setelah menampilkan semua rumah daerah
	}
}
