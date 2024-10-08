package main

import (
	"fmt"
)

func main() {
	// Membaca 5 data integer
	var a, b, c, d, e int
	fmt.Println("Masukkan 5 angka integer (nilai antara 32 s.d. 127):")
	fmt.Scanf("%d %d %d %d %d\n", &a, &b, &c, &d, &e)

	// Mencetak representasi karakter dari data integer
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	// Membaca 3 data karakter tanpa spasi
	var x, y, z rune
	fmt.Println("Masukkan 3 karakter tanpa spasi:")
	fmt.Scanf("%c%c%c", &x, &y, &z)

	// Mencetak karakter setelah karakter input (berdasarkan tabel ASCII)
	fmt.Printf("%c%c%c\n", x+1, y+1, z+1)
}
