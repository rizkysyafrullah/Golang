package main

import "fmt"

func main() {
	var point = 9

	if point >= 7 {
		switch point {
		case 10:
			fmt.Println("Sempurna")
		default:
			fmt.Println("Luar Biasa")
		}
	} else {
		if point == 5 {
			fmt.Println("Lumayan")
		} else if point >= 3 {
			fmt.Println("Kurang")
		} else {
			fmt.Println("Sangat Kurang")
		}
	}
}
