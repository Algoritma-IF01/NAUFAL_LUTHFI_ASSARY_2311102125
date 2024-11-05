package main

import (
	"fmt"
)

// Struct untuk menyimpan data klub
type Klub_125 struct {
	Nama string
}

// Fungsi untuk menambah skor
func tambahSkor_125(skorA int, skorB int) string {
	if skorA > skorB {
		return "A"
	} else if skorA < skorB {
		return "B"
	} else {
		return "Draw"
	}
}

func main() {
	var klubA, klubB Klub_125
	var skorA, skorB int
	var hasil string
	pemenang_125 := make([]string, 0)

	// Input klub A dan B
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA.Nama)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB.Nama)

	// Loop untuk input skor pertandingan
	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d (skor %s vs %s): ", i, klubA.Nama, klubB.Nama)
		fmt.Scanf("%d %d\n", &skorA, &skorB)

		// Berhenti jika skor tidak valid (negatif)
		if skorA < 0 || skorB < 0 {
			break
		}

		// Tentukan pemenang pertandingan
		hasil = tambahSkor_125(skorA, skorB)
		if hasil == "A" {
			pemenang_125 = append(pemenang_125, klubA.Nama)
		} else if hasil == "B" {
			pemenang_125 = append(pemenang_125, klubB.Nama)
		} else {
			pemenang_125 = append(pemenang_125, "Draw")
		}
	}

	// Cetak hasil akhir
	fmt.Println("\nHasil Pertandingan:")
	for i, klub := range pemenang_125 {
		fmt.Printf("Hasil %d: %s\n", i+1, klub)
	}
	fmt.Println("Pertandingan selesai")
}
