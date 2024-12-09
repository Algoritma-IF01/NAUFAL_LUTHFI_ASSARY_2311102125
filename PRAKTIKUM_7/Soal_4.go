package main

import (
	"bufio"   // Untuk membaca input secara baris per baris
	"fmt"     // Untuk menampilkan output ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk mengkonversi string ke integer
	"strings" // Untuk memanipulasi string
)

const maxBooks = 7919 

type Buku struct {
	id        int    
	judul     string 
	penulis   string 
	penerbit  string 
	tahun     int    
	rating    int    
}

// Fungsi untuk mendaftarkan buku ke dalam perpustakaan
func DaftarkanBuku(perpustakaan []Buku, jumlahBuku int) []Buku {
	reader := bufio.NewReader(os.Stdin) // Membuat reader untuk menerima input
	for i := 0; i < jumlahBuku; i++ {
		fmt.Printf("\nMasukkan data untuk buku ke-%d:\n", i+1)

		// Meminta ID buku
		fmt.Print("ID: ")
		idInput, _ := reader.ReadString('\n')
		idInput = strings.TrimSpace(idInput)
		id, _ := strconv.Atoi(idInput)

		// Meminta Judul buku
		fmt.Print("Judul: ")
		judul, _ := reader.ReadString('\n')
		judul = strings.TrimSpace(judul)

		// Meminta Penulis buku
		fmt.Print("Penulis: ")
		penulis, _ := reader.ReadString('\n')
		penulis = strings.TrimSpace(penulis)

		// Meminta Penerbit buku
		fmt.Print("Penerbit: ")
		penerbit, _ := reader.ReadString('\n')
		penerbit = strings.TrimSpace(penerbit)

		// Meminta Tahun terbit buku
		fmt.Print("Tahun: ")
		tahunInput, _ := reader.ReadString('\n')
		tahunInput = strings.TrimSpace(tahunInput)
		tahun, _ := strconv.Atoi(tahunInput)

		// Meminta Rating buku
		fmt.Print("Rating: ")
		ratingInput, _ := reader.ReadString('\n')
		ratingInput = strings.TrimSpace(ratingInput)
		rating, _ := strconv.Atoi(ratingInput)

		// Menyimpan buku ke dalam perpustakaan
		perpustakaan[i] = Buku{id, judul, penulis, penerbit, tahun, rating}
	}
	return perpustakaan
}

// Fungsi untuk mengurutkan buku berdasarkan rating secara menurun menggunakan metode Insertion Sort
func UrutkanBuku(perpustakaan []Buku, jumlahBuku int) {
	for i := 1; i < jumlahBuku; i++ {
		key := perpustakaan[i]  // Buku yang akan diposisikan
		j := i - 1

		// Memindahkan buku dengan rating lebih rendah ke kanan
		for j >= 0 && perpustakaan[j].rating < key.rating {
			perpustakaan[j+1] = perpustakaan[j]
			j--
		}
		perpustakaan[j+1] = key // Menempatkan buku pada posisi yang tepat
	}
}

// Fungsi untuk menampilkan buku dengan rating tertinggi
func TampilkanBukuFavorit(perpustakaan []Buku, jumlahBuku int) {
	if jumlahBuku == 0 { // Jika perpustakaan kosong
		fmt.Println("Tidak ada buku di perpustakaan.") 
		return
	}

	favoriteIdx := 0                             // Indeks buku dengan rating tertinggi
	highestRating := perpustakaan[0].rating      // Rating dari buku pertama

	// Menelusuri buku untuk mencari yang memiliki rating tertinggi
	for i := 1; i < jumlahBuku; i++ {
		if perpustakaan[i].rating > highestRating {
			highestRating = perpustakaan[i].rating
			favoriteIdx = i
		}
	}

	// Menampilkan buku dengan rating tertinggi
	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n",
		perpustakaan[favoriteIdx].judul, perpustakaan[favoriteIdx].penulis,
		perpustakaan[favoriteIdx].penerbit, perpustakaan[favoriteIdx].tahun,
		perpustakaan[favoriteIdx].rating)
}

// Fungsi untuk menampilkan lima buku teratas berdasarkan rating tertinggi
func Tampilkan5BukuTeratas(perpustakaan []Buku, jumlahBuku int) {
	fmt.Println("\n5 Buku Teratas dengan Rating Tertinggi:")
	// Memastikan buku sudah diurutkan sebelum ditampilkan
	UrutkanBuku(perpustakaan, jumlahBuku)
	for i := 0; i < 5 && i < jumlahBuku; i++ {
		fmt.Printf("Judul: %s, Rating: %d\n", perpustakaan[i].judul, perpustakaan[i].rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating tertentu
func CariBukuBerdasarkanRating(perpustakaan []Buku, jumlahBuku int, ratingYangDicari int) {
	found := false // Menandakan apakah buku ditemukan

	// Mencari buku dengan rating tertentu
	fmt.Printf("\nBuku dengan rating %d:\n", ratingYangDicari)
	for i := 0; i < jumlahBuku; i++ {
		if perpustakaan[i].rating == ratingYangDicari {
			// Menampilkan detail buku
			fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n\n",
				perpustakaan[i].judul, perpustakaan[i].penulis,
				perpustakaan[i].penerbit, perpustakaan[i].tahun,
				perpustakaan[i].rating)
			found = true
		}
	}

	// Jika tidak ada buku dengan rating tersebut
	if !found {
		fmt.Printf("Tidak ada buku dengan rating %d\n", ratingYangDicari)
	}
}

// Program utama
func main() {
	reader := bufio.NewReader(os.Stdin) // Membuat reader untuk menerima input

	// Meminta jumlah buku yang ada di perpustakaan
	fmt.Print("Masukkan jumlah buku di perpustakaan: ")
	jumlahBukuInput, _ := reader.ReadString('\n')
	jumlahBukuInput = strings.TrimSpace(jumlahBukuInput)
	jumlahBuku, _ := strconv.Atoi(jumlahBukuInput)

	// Membuat pustaka dengan jumlah buku yang ditentukan
	var perpustakaan = make([]Buku, jumlahBuku)

	// Mendaftarkan buku ke dalam perpustakaan
	perpustakaan = DaftarkanBuku(perpustakaan, jumlahBuku)

	// Mengurutkan buku berdasarkan rating
	UrutkanBuku(perpustakaan, jumlahBuku)

	// Menampilkan buku dengan rating tertinggi
	TampilkanBukuFavorit(perpustakaan, jumlahBuku)

	// Menampilkan lima buku teratas berdasarkan rating tertinggi
	Tampilkan5BukuTeratas(perpustakaan, jumlahBuku)

	// Mencari buku berdasarkan rating yang diinginkan
	var ratingYangDicari int
	fmt.Print("\nMasukkan rating untuk mencari buku: ")
	fmt.Scanf("%d\n", &ratingYangDicari)
	CariBukuBerdasarkanRating(perpustakaan, jumlahBuku, ratingYangDicari)
}
