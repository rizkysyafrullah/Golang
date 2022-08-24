package main

import "fmt"

func main() {
	var point = 1
	if point == 10 {
		fmt.Println("Anda Lulus dengan skor SEMPURNA!, skor anda =", point)
	} else if point >= 5 {
		fmt.Println("Anda Lulus !, skor anda =", point)
	} else if point == 4 {
		fmt.Println("Anda Hampir Lulus !, skor anda =", point)
	} else {
		fmt.Println("Anda tidak lulus, nilai anda =", point)
	}
}
