package main

import "fmt"

func main() {
	var point = 2

	switch {
	case point == 8:
		fmt.Println("Sempurna")
	case (point < 8) && (point > 3):
		fmt.Println("Bagus")
	default:
		{
			fmt.Println("Yaah lumayan lah")
			fmt.Println("Belajar lagi Yaa")
		}
	}
}
