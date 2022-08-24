package main

import "fmt"

func main() {
	var buah = [4]string{"Duren", "Apel", "Pir", "Jeruk"}

	for i := 0; i < len(buah); i++ {
		fmt.Println("elemen", i, ":", buah[i])
	}
}
