package main

import "fmt"

func main() {
	dst := make([]string, 3)
	src := []string{"pineapple", "starfruit", "banana", "grape"}
	n := copy(dst, src) //outputnya nilai jumlah objek yg di copy, yg di copy sebanyak len(dst)

	fmt.Println(" ")
	fmt.Println("source :", src)      //[pineapple starfruit banana grape]
	fmt.Println("destination :", dst) //[pineapple starfruit]
	fmt.Println(" ")
	fmt.Println("jumlah object yg di copy :", n) //3
	fmt.Println(" ")
}
