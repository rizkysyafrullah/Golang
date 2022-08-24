package main

import "fmt"

func main() {
	var buah = [4]string{"mangga", "srikaya", "apel", "jeruk"}

	/*for i, buah := range buah {
		fmt.Printf("nama buah : %s\n", buah)
	}*/

	/*aturan di golang tidak boleh ada variable yg ga kepake
	outputnya error karena si var i ga dipake
	solusinya pake variable underscore (_) niar ditampung*/

	/*for _, buah := range buah {
		fmt.Printf("nama buah : %s\n", buah)
	}*/

	// sdngkn kalo butuh indexnya doang, maka buah nya yg di tampung
	for i, _ := range buah {
		fmt.Println("nama buah :", i)
	}

	// atau

	/*for i := range buah {
		fmt.Println("nama buah :", i)
	}*/
}
