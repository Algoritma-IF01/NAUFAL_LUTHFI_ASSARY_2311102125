package main

import (
	"fmt"
)

// Definisi fungsi
func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

// Fungsi komposisi
func fogoh(x int) int {
	return f(g(h(x)))
}

func gohof(x int) int {
	return g(h(f(x)))
}

func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	// Input
	var a, b, c int
	fmt.Print("Masukkan tiga bilangan (a, b, c) dipisahkan oleh spasi: ")
	fmt.Scan(&a, &b, &c)

	// Output
	fmt.Println("fogoh(a) =", fogoh(a))
	fmt.Println("gohof(b) =", gohof(b))
	fmt.Println("hofog(c) =", hofog(c))
}
