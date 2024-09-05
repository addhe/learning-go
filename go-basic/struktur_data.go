package main

import "fmt"

func main() {
	var bilangan [5]int = [5]int{10, 20, 30, 40, 50}

	fmt.Println("Elemen pertama:", bilangan[0])
	fmt.Println("Elemen kedua:", bilangan[1])

	// mari kita loop
	fmt.Println("Semua elemen pada variable bilangan:")
	for i := 0; i < len(bilangan); i++ {
		fmt.Println(bilangan[i])
	}
}
