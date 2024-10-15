package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 3; i++ {
		var beratParsel int

		// Meminta input berat parsel dalam gram dari pengguna
		fmt.Printf("Berat parsel #%d (gram): ", i)
		fmt.Scan(&beratParsel)

		// Menghitung total berat dalam kilogram dan sisa gramnya
		totalKg := beratParsel / 1000
		sisaGram := beratParsel % 1000
		// Biaya per kilogram adalah Rp. 10.000
		biayaPerKg := 10000
		biayaTambahan := 0

		// Menghitung biaya tambahan berdasarkan sisa gram
		if sisaGram > 0 {
			if sisaGram <= 500 {
				// Jika sisa gram kurang dari atau sama dengan 500 gram, biaya tambahan Rp. 5 per gram
				biayaTambahan = sisaGram * 5
			} else {
				// Jika sisa gram lebih dari 500 gram, biaya tambahan Rp. 15 per gram
				biayaTambahan = sisaGram * 15
			}
		}

		// Menghitung total biaya
		totalBiaya := (totalKg * biayaPerKg) + biayaTambahan

		// Jika total berat lebih dari atau sama dengan 10 kg, biaya sisa gram dihapuskan
		if totalKg >= 10 {
			biayaTambahan = 0
			totalBiaya = totalKg * biayaPerKg
		}

		// Menampilkan detail berat dan biaya
		fmt.Printf("Detail berat: %d kg + %d gram\n", totalKg, sisaGram)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", totalKg*biayaPerKg, biayaTambahan)
		fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
		fmt.Println() // Untuk memisahkan setiap output program
	}
}