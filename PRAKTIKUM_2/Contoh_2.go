package main 
import "fmt"

func main(){
	var a, b, c float64
	var hipotenusa bool
	fmt.Print("Masukkan Nilai A")
	fmt.Scanln(&a)
	fmt.Print("Masukkan Nilai B")
	fmt.Scanln(&b)
	fmt.Print("Masukkan Nilai C")
	fmt.Scanln(&c)
	hipotenusa = (c*c) == (a*a + b*b)
	fmt.Println( "Sisi c adalah hipotenusa a,b,c: ", hipotenusa)
}