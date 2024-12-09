package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan metode insertion sort
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ { // Mulai dari elemen kedua
		key := arr[i]  // Ambil elemen yang akan disisipkan
		j := i - 1     // Tentukan posisi untuk pengecekan elemen sebelumnya

		// Geser elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]  // Geser elemen
			j--  // Mundur ke elemen sebelumnya
		}
		arr[j+1] = key  // Sisipkan elemen key pada posisi yang benar
	}
	return arr  // Kembalikan array yang sudah terurut
}

// Fungsi untuk memeriksa apakah jarak antar elemen array teratur
func checkDistance(arr []int) string {
	if len(arr) <= 1 {  // Jika array kosong atau hanya memiliki satu elemen
		return "Data berjarak tidak tetap"  // Tidak ada jarak antar elemen
	}

	distance := arr[1] - arr[0]  // Hitung jarak antar elemen pertama dan kedua
	for i := 1; i < len(arr)-1; i++ {  // Periksa semua elemen
		if arr[i+1]-arr[i] != distance {  // Jika jaraknya tidak tetap
			return "Data berjarak tidak tetap"
		}
	}
	return fmt.Sprintf("Data berjarak %d", distance)  // Jika jarak antar elemen tetap
}

func main() {
	var input int
	var data []int  // Array untuk menyimpan input pengguna

	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")
	for {
		fmt.Scan(&input)  // Menerima input angka
		if input < 0 {  // Jika input negatif, berhenti
			break
		}
		data = append(data, input)  // Menambahkan input ke dalam array
	}

	if len(data) == 0 {  // Jika tidak ada data yang dimasukkan
		fmt.Println("Tidak ada data untuk diproses.")
		return
	}

	// Urutkan array menggunakan insertion sort
	sortedData := insertionSort(data)

	// Periksa apakah jarak antar elemen teratur
	status := checkDistance(sortedData)

	// Menampilkan hasil
	fmt.Println("Keluaran:")
	for _, v := range sortedData {
		fmt.Printf("%d ", v)  // Tampilkan hasil urutan
	}
	fmt.Println()
	fmt.Println(status)  // Tampilkan status jarak antar elemen
}
