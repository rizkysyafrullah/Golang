package main

import "fmt"

func main() {
	/*var numbers1 = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	var numbers2 = [2][3]int{{3, 2, 1}, {6, 5, 4}} //mending ini lebih simple wkwk

	fmt.Println("Numbers1 =", numbers1)
	fmt.Println("Numbers2 =", numbers2)*/

	var char1 = [3][4]string{{"A", "B", "C", "D"}, {"E", "F", "G", "H"}, {"I", "J", "K", "L"}}
	var char2 = [4][4]string{{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"U", "V", "W", "X"},
		{"Y", "Z", "Z", "Z"}}

	fmt.Println("Char 1 :", char1)
	fmt.Println("Char 2 :", char2)
}
