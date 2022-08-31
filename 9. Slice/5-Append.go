package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits, "papaya") // nambahin 'papaya' ke variabel fruits

	fmt.Println(fruits)  //[apple grape banana]
	fmt.Println(cFruits) //[apple grape banana papaya]
}
