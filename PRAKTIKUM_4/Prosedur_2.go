package main

import (
    "bufio"   
    "fmt"     
    "os"      // digunakan untuk akses ke fungsi sistem operasi, seperti membaca input dari standar input
    "strings" // digunakan untuk memanipulasi string, seperti memisahkan input menjadi bagian-bagian
    "strconv" // digunakan untuk mengubah string menjadi tipe data lain, seperti integer
)

// Prosedur untuk menghitung skor dan total waktu dari soal yang dikerjakan peserta
func hitungSkor(soal []int) (int, int) {
    totalSkor := 0
    totalWaktu := 0

    for _, waktu := range soal {
        if waktu <= 300 {
            totalSkor++
            totalWaktu += waktu
        } else {
            totalWaktu += 0
        }
    }

    return totalSkor, totalWaktu
}

// Fungsi utama untuk membaca input, memproses data peserta, dan menentukan pemenang
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var pemenang string
    maxSkor := -1
    minWaktu := 999999
    var skor, totalWaktu int

    fmt.Println("Masukkan data peserta, akhiri dengan 'Selesai':")
    for scanner.Scan() {
        input := scanner.Text()
        if strings.ToLower(input) == "selesai" {
            break
        }

        // Memisahkan input menjadi nama peserta dan waktu pengerjaan soal
        data := strings.Fields(input)
        if len(data) != 9 {
            fmt.Println("Input tidak valid, pastikan formatnya: nama waktu1 waktu2 ... waktu8")
            continue
        }

        peserta := data[0]
        soal := make([]int, 8)
        valid := true

        // Mengubah waktu pengerjaan dari string ke integer
        for i := 1; i <= 8; i++ {
            waktu, err := strconv.Atoi(data[i])
            if err != nil {
                fmt.Println("Waktu harus berupa angka.")
                valid = false
                break
            }
            soal[i-1] = waktu
        }

        if !valid {
            continue
        }

        // Menghitung skor dan total waktu untuk peserta saat ini
        skor, totalWaktu = hitungSkor(soal)

        // Tentukan pemenang berdasarkan jumlah soal yang diselesaikan dan waktu total
        if skor > maxSkor || (skor == maxSkor && totalWaktu < minWaktu) {
            maxSkor = skor
            minWaktu = totalWaktu
            pemenang = peserta
        }
    }

    // Mencetak hasil pemenang
    if pemenang != "" {
        fmt.Printf("%s %d %d\n", pemenang, maxSkor, minWaktu)
    } else {
        fmt.Println("Tidak ada peserta yang valid.")
    }
}