package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n)

	esTebak := 0
	esCendol := 0
	lamang := 0

	for i := 0; i < n; i++ {
		var cardNumber int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i+1)
		fmt.Scan(&cardNumber)

		cardStr := strconv.Itoa(cardNumber)
		allOdd := true
		allEven := true

		for _, digit := range cardStr {
			intDigit := int(digit - '0')
			if intDigit%2 != 1 {
				allOdd = false
			}
			if intDigit%2 != 0 {
				allEven = false
			}
		}

		if allOdd {
			fmt.Println("Es Tebak")
			esTebak++
		} else if allEven {
			fmt.Println("Es Cendol")
			esCendol++
		} else {
			fmt.Println("Lamang")
			lamang++
		}
	}

	fmt.Printf("Jumlah peserta yang memperoleh Es Tebak: %d\n", esTebak)
	fmt.Printf("Jumlah peserta yang memperoleh Es Cendol: %d\n", esCendol)
	fmt.Printf("Jumlah peserta yang memperoleh Lamang: %d\n", lamang)
}
