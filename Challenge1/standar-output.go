package main

import (
	"fmt"
)

func main() {

	// menampilkan tipe data dari variabel i
	i := 21
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	// menampilkan tanda %
	fmt.Printf("%% \n")

	// menampilkan nilai boolean j : true
	var j bool = true
	fmt.Printf("%t\n \n", j)

	// menampilkan unicode russia : Я (ya)
	unicodeRusia := '\u042F'
	fmt.Printf("%c\n", unicodeRusia)

	// menampilkan nilai base 10 : 21
	fmt.Printf("%d \n", i)

	// menampilkan nilai base 8 :25
	fmt.Printf("%o \n", i)

	// menampilkan nilai base 16 : f
	fmt.Printf("%x \n", 15)

	// menampilkan nilai base 16 : F
	fmt.Printf("%X \n", 15)

	// menampilkan unicode karakter Я : U+042F var k float64 = 123.456;
	var charRusia = 'Я'
	fmt.Printf("%U \n \n", charRusia)

	// menampilkan float : 123.456000
	var k float64 = 123.456
	fmt.Printf("%f\n", k)

	// menampilkan float scientific : 1.234560E+02
	fmt.Printf("%e", k)

}
