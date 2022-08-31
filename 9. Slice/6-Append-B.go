package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana"}
	var bFruits = fruits[0:2]

	fmt.Println(cap(bFruits)) // 3
	fmt.Println(len(bFruits)) // 2
	fmt.Println(" ")
	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(bFruits) // ["apple", "grape"]
	fmt.Println(" ")

	var cFruits = append(bFruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	fmt.Println(bFruits) // ["apple", "grape"]
	fmt.Println(cFruits) // ["apple", "grape", "papaya"]
	fmt.Println(" ")
}
