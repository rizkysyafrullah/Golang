/*
Switch merupakan seleksi kondisi yang sifatnya fokus pada satu variabel,
lalu kemudian di-cek nilainya. Contoh sederhananya seperti penentuan
apakah nilai variabel x adalah: 1 , 2 , 3 , atau lainnya.
*/
package main

import "fmt"

func main() {
	var point = 7

	switch point {
	case 8:
		fmt.Println("SEMPURNA!")
	case 7:
		fmt.Println("Bagus!")
	default: //semacam 'else' pada sebuah switch
		fmt.Println("Not Bad")
	}
}
