package main

import (
    "fmt"
    "math"
)

// Mengisi array berdasarkan input pengguna
func isiArray_125(arr []int) {
    for i := range arr {
        fmt.Printf("Masukkan elemen ke-%d: ", i)
        fmt.Scan(&arr[i])
    }
}

// Menampilkan keseluruhan isi dari array
func tampilKeseluruhan_125(arr []int) {
    fmt.Println("a. Isi array:", arr)
}

// Menampilkan elemen dengan indeks ganjil saja (dimulai dari indeks 1, 3, 5, dst)
func tampilIndeksGanjil_125(arr []int) {
    fmt.Printf("b. Elemen pada indeks ganjil: ")
    for i := 1; i < len(arr); i += 2 {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

// Menampilkan elemen dengan indeks genap saja (dimulai dari indeks 0, 2, 4, dst)
func tampilIndeksGenap_125(arr []int) {
    fmt.Printf("c. Elemen pada indeks genap: ")
    for i := 0; i < len(arr); i += 2 {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

// Menampilkan elemen array dengan indeks kelipatan bilangan tertentu
func tampilKelipatanIndeks_125(arr []int, x int) {
    if x <= 0 {
        fmt.Println("Nilai kelipatan harus lebih besar dari 0")
        return
    }
    fmt.Printf("   Elemen pada indeks kelipatan %d: ", x)
    for i := x; i < len(arr); i += x {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

// Menghapus elemen pada indeks tertentu
func hapusIndeks_125(arr []int, index int) []int {
    if index < 0 || index >= len(arr) {
        fmt.Println("Indeks tidak valid")
        return arr
    }
    fmt.Printf("   Menghapus elemen pada indeks %d\n", index)
    return append(arr[:index], arr[index+1:]...)
}

// Menghitung rata-rata dari elemen array
func rataRataArray_125(arr []int) float64 {
    total := 0
    for _, v := range arr {
        total += v
    }
    return float64(total) / float64(len(arr))
}

// Menghitung standar deviasi dari elemen array
func standarDeviasiArray_125(arr []int) float64 {
    rataRata := rataRataArray_125(arr)
    var sum float64
    for _, v := range arr {
        sum += math.Pow(float64(v)-rataRata, 2)
    }
    variance := sum / float64(len(arr))
    return math.Sqrt(variance)
}

// Menghitung frekuensi dari suatu bilangan di dalam array
func frekuensiBilangan_125(arr []int, nilai int) int {
    count := 0
    for _, v := range arr {
        if v == nilai {
            count++
        }
    }
    return count
}

func main() {
    // Meminta jumlah elemen dari pengguna
    var N int
    fmt.Print("Masukkan jumlah elemen array (N): ")
    fmt.Scan(&N)

    // Membuat array dan mengisi dengan input pengguna
    arr := make([]int, N)
    isiArray_125(arr)

    // a. Menampilkan isi array
    tampilKeseluruhan_125(arr)

    // b. Menampilkan elemen pada indeks ganjil
    tampilIndeksGanjil_125(arr)

    // c. Menampilkan elemen pada indeks genap
    tampilIndeksGenap_125(arr)

    // d. Menampilkan elemen pada indeks kelipatan bilangan tertentu (misal x = 3)
    var x int
    fmt.Print("d. Masukkan nilai kelipatan (X): ")
    fmt.Scan(&x)
    tampilKelipatanIndeks_125(arr, x)

    // e. Menghapus elemen pada indeks tertentu (misal indeks yang ingin dihapus adalah 4)
    var index int
    fmt.Print("e. Masukkan indeks yang ingin dihapus: ")
    fmt.Scan(&index)
    arr = hapusIndeks_125(arr, index)
    fmt.Println("   Array setelah dihapus: ",arr)

    // f. Menghitung dan menampilkan rata-rata
    rataRata := rataRataArray_125(arr)
    fmt.Printf("f. Rata-rata dari elemen array: %.2f\n", rataRata)

    // g. Menghitung dan menampilkan standar deviasi
    stdDev := standarDeviasiArray_125(arr)
    fmt.Printf("g. Standar deviasi dari elemen array: %.2f\n", stdDev)

    // h. Menghitung frekuensi dari suatu bilangan (misal nilai yang dicari adalah 3)
    var nilai int
    fmt.Print("h. Masukkan nilai yang ingin dicari frekuensinya: ")
    fmt.Scan(&nilai)
    frekuensi := frekuensiBilangan_125(arr, nilai)
    fmt.Printf("   Frekuensi bilangan %d dalam array: %d\n", nilai, frekuensi)
}
