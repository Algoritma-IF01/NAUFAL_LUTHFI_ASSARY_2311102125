package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Membaca input dari pengguna
	fmt.Println("Masukkan jumlah daerah:")
	scanner.Scan() // Menerima input jumlah daerah
	t, _ := strconv.Atoi(scanner.Text()) // Mengonversi input menjadi integer

	// Loop untuk menerima data setiap daerah
	for i := 0; i < t; i++ {
		fmt.Printf("Masukkan angka untuk daerah %d (pisahkan dengan spasi):\n", i+1)
		scanner.Scan() // Menerima input deretan angka
		input := scanner.Text()

		// Memanggil fungsi untuk mengubah input menjadi array bilangan bulat
		nums := parseInput(input)

		// Memisahkan angka menjadi ganjil dan genap
		odds, evens := separateOddEven(nums)

		// Mengurutkan bilangan ganjil secara menaik dan bilangan genap secara menurun
		sort.Ints(odds)
		sort.Sort(sort.Reverse(sort.IntSlice(evens)))

		// Menampilkan hasil urutan bilangan ganjil dan genap
		fmt.Printf("Hasil daerah %d:\n", i+1)
		printResult(odds, evens)
	}
}

// Fungsi untuk mengonversi input string menjadi array bilangan bulat
func parseInput(input string) []int {
	parts := strings.Fields(input) // Memisahkan input berdasarkan spasi
	nums := make([]int, len(parts)) // Membuat array untuk menampung angka
	for i, p := range parts {
		nums[i], _ = strconv.Atoi(p) // Mengonversi string menjadi angka
	}
	return nums
}

// Fungsi untuk memisahkan bilangan ganjil dan genap
func separateOddEven(nums []int) ([]int, []int) {
	var odds, evens []int // Membuat slice untuk bilangan ganjil dan genap
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num) // Jika genap, tambahkan ke slice evens
		} else {
			odds = append(odds, num) // Jika ganjil, tambahkan ke slice odds
		}
	}
	return odds, evens 
}

// Fungsi untuk menampilkan hasil setelah proses pemisahan dan pengurutan
func printResult(odds []int, evens []int) {
	for _, odd := range odds { // Menampilkan bilangan ganjil
		fmt.Print(odd, " ")
	}
	for _, even := range evens { // Menampilkan bilangan genap
		fmt.Print(even, " ")
	}
	fmt.Println() // Pindah baris setelah hasil ditampilkan
}
