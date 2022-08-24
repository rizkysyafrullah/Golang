package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "banana", "watermelon", "grape"}

	//bisa juga ditulis scr vertikal ygy
	/*var fruits = [4]string{
	"apple",
	"banana",
	"watermelon",
	"grape"}*/

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element buah \t", fruits)
}
