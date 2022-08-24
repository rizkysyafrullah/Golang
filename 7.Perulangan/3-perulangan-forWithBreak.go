package main

import "fmt"

func main() {
	w := 0

	for w < 5 {
		fmt.Println("Nomor", w)

		w++
		if w == 5 {
			break
		}
	}
}

/*
Dalam perulangan tanpa henti di atas, variabel i yang nilai awalnya 0 di-inkrementasi. Ketika nilai i sudah mencapai 5 ,
keyword break digunakan, dan perulangan akan berhenti.
*/
