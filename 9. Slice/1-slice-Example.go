package main

import "fmt"

func main() {
	/*var buah2an = []string{"Pisang", "Apel", "Durian", "Salak"}
	fmt.Println(buah2an[3]) //salak

	/*var angka = []int{10, 20, 30, 40}
	fmt.Println(angka[3])*/ //40

	var buah2an = []string{"Pisang", "Apel", "Durian", "Salak"}
	var sumBuah = buah2an[0:3] // [Pisang Apel Durian]
	/*var sumBuah = buah2an[0:2] // [Pisang Apel]
	var sumBuah = buah2an[0:0] // kosong
	var sumBuah = buah2an[4:0] // error --> misal buah2an[a:b] --> maka wajib nilai a < b
	var sumBuah = buah2an[:] // semua
	var sumBuah = buah2an[2:] //[Durian Salak]
	var sumBuah = buah2an[:2] // [Pisang Apel]*/
	fmt.Println(sumBuah)

}
