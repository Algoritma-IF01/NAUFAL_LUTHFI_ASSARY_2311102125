package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// Fungsi untuk mengecek apakah titik berada di dalam lingkaran
func dalamLingkaran(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	for i := 1; i <= 4; i++ {
	// Input koordinat pusat dan radius untuk dua lingkaran
	var cx1, cy1, r1, cx2, cy2, r2 float64
	fmt.Println("Masukkan bilangan: ")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)

	// Input koordinat titik yang akan dicek
	var x, y float64
	fmt.Scan(&x, &y)

	// Mengecek posisi titik terhadap kedua lingkaran
	inCircle1 := dalamLingkaran(cx1, cy1, r1, x, y)
	inCircle2 := dalamLingkaran(cx2, cy2, r2, x, y)

	// Menentukan keluaran berdasarkan posisi titik
	if inCircle1 && inCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
}
