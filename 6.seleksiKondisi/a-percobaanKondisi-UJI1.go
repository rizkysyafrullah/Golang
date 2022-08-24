package main

import "fmt"

func main() {
	var mtk = 81
	var ipa = 81

	var niltot = (mtk + ipa) / 2

	if niltot >= 85 {
		fmt.Println("Sangat Baik, nilai anda", niltot)
	} else if niltot >= 75 {
		fmt.Println("Baik, nilai anda", niltot)
	} else if niltot >= 60 {
		fmt.Println("Kurang, nilai anda", niltot)
	} else {
		fmt.Println("Sangat Kurang, nilai anda", niltot)
	}
}
