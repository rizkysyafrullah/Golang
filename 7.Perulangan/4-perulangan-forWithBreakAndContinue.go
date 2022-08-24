package main

func main() {
	/*for s := 1; s <= 1-0; s++ {
		if s%2 == 1 { //kalo ganjil di kontinyu mksdny
			continue
		}

		if s > 80 {
			break
		}

		println("Nomor", s)
	}*/

	for z := 1; z <= 100; z++ {
		if z%2 == 0 { //kalo genap di kontinyu mksdny
			continue
		}

		if z > 80 {
			break
		}

		println("Nomor", z)
	}
}
