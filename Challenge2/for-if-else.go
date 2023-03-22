package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
	}

	for j := 0; j < 5; j++ {
		fmt.Println("Nilai j =", j)
	}

	for bytePosition, character := range "ɋАШАɊВО" {
		fmt.Printf("character %#U start at byte position %d \n", character, bytePosition)
	}

	for j := 5; j < 100; j++ {
		fmt.Println("Nilai j =", j)
		if j == 10 {
			break
		}
	}
}
