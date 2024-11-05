package main

import (
	"bufio"   
	"fmt"     
	"os"      
	"strings" 
)

const NMAX int = 127 

type tabel [NMAX]rune 
/*I.S. Data tersedia dalam piranti masukan
  F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
  Proses input selama karakter bukanlah TITIK dan n <= NMAX*/
func isiArray(t *tabel, n *int, line string) {
	*n = 0                      
	for _, char := range line { 
		if *n >= NMAX { 
			break 
		}
		t[*n] = char 
		*n++         
	}
}

/*I.S. Terdefinisi array t yang berisi sejumlah n karakter
  F.S. n karakter dalam array muncul di layar*/
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ { 
		fmt.Print(string(t[i]), " ") 
	}
	fmt.Println() 
}

/*I.S. Terdefinisi array t yang berisi sejumlah n karakter
  F.S. Urutan isi array t terbalik
*/
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ { 
		t[i], t[n-1-i] = t[n-1-i], t[i] 
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ { 
		if t[i] != t[n-1-i] { 
			return false 
		}
	}
	return true 
}

func main() {
	// si array tab dengan memanggil prosedur isiArray
	var tab tabel                         
	var m int                             
	scanner := bufio.NewScanner(os.Stdin) 
	fmt.Println("Masukkan teks (ketik '.' untuk berhenti): ")

	for scanner.Scan() { 
		line := scanner.Text()            
		if strings.ToUpper(line) == "." { 
			break 
		}
		isiArray(&tab, &m, line) 
		
		fmt.Print("Teks         : ")
		cetakArray(tab, m) 

		// Balikan isi array tab dengan memanggil balkanArray
		balikanArray(&tab, m) 
		fmt.Print("Reverse Teks : ")
		cetakArray(tab, m) 

		// Mengecek dan menampilkan hasil palindrom
		isPalindrom := palindrom(tab, m)           
		fmt.Println("Palindrom    :", isPalindrom) 
		fmt.Println()                              
	}
}
