package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan nilai m: ")
	fmt.Scan(&m)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" + ")
		}
		fmt.Print(m)
	}

	hasil := multiply(n, m)
	fmt.Printf(" = %d\n", hasil)
}

func multiply(n, m int) int {
	if m == 0 {
		return 0
	}
	if m < 0 {
		return -multiply(n, -m)
	}
	return n + multiply(n, m-1)
}
