package main

import "fmt"

//import "strconv"

func main() {
	var i int = 21
	var tandaPersen int = 37
	var j bool = true
	fmt.Printf("%d \n%T\n", i, i)   // Nilai variabel i dan type data
	fmt.Printf("%c\n", tandaPersen) // Mencetak karakter
	fmt.Printf("%t\n", j)           //Mengambil nilai dari tipe data boolean

	fmt.Printf("\n")
	var biner int = 21
	var tandaTanya = 63
	var cetakBase10 int = 21
	var cetakBase16 int = 15
	var unicode int32 = 'Я'

	fmt.Printf("%b\n", biner)       //Mencetak(konversi) nilai ke biner
	fmt.Printf("%c\n", tandaTanya)  // Mencetak karakter
	fmt.Printf("%d\n", cetakBase10) // Mencetak nilai base 10 dengan menggunakan %d
	fmt.Printf("%o\n", cetakBase10) // Mencetak nilai base 8 dengan menggunaan %o
	fmt.Printf("%x\n", cetakBase16) // Mencetak nilai base 16 ke dalam lower case menggunakan %x
	fmt.Printf("%X\n", cetakBase16) // Mencetak nilai base 16 ke dalam upper case menggunakan %X
	fmt.Printf("%U\n", unicode)     // Mencetak nilai karakter unicode

	var k float64 = 123.456
	fmt.Printf("\n%3f\n", k) // menampilkan float
	fmt.Printf("%E", k)      // menampilkan float scientific

}
