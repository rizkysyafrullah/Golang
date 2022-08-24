package main

//sama kaya 1-perulangan-for.go , cuma beda style aja
/* dengan menuliskan kondisi setelah keyword for (hanya kondisi). Deklarasi dan iterasi variabel counter tidak
dituliskan setelah keyword, hanya kondisi perulangan saja. Konsepnya mirip seperti while milik bahasa pemrograman lain.*/

import "fmt"

func main() {
	var k = 0

	for k < 5 {
		fmt.Println("Nomor", k)
		k++
	}
}
