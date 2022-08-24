package main

import "fmt"

func main() {
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		//Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang menggunakan tanda :=
		fmt.Printf("perfect!", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}
