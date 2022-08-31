package main

import "fmt"

func main() {
	dst := []string{"a", "b", "c"}
	src := []string{"A", "B"}
	result := copy(dst, src)

	fmt.Println(" ")
	fmt.Println(src)    //[A B]
	fmt.Println(dst)    //[A B c]
	fmt.Println(result) //2
	fmt.Println(" ")
}
