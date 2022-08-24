package main

import "fmt"

func main() {
	for i := 0; i < 7; i++ {
		for j := i; j < 7; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
