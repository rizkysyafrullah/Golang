package main

import "fmt"

func main() {
	var point = 7

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("Bagus")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
		//fallthrough
	default:
		{ //kalo mau kebaca line ini maka harus nyalai fallthrough di line 16
			fmt.Println("Not bad")
			fmt.Println("you need to learn more")
		}
	}
}
