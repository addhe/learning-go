package main

import "fmt"

func main() {
	var buah []string

	buah = append(buah, "Apel")
	buah = append(buah, "Mangga", "Jeruk")

	fmt.Println("Buah pertama:", buah[0])
	fmt.Println("Buah kedua:", buah[1])
	fmt.Println("Buah ketiga:", buah[2])

	fmt.Println("Semua Buah:", buah)

	buahFavorite := buah[1:3]
	fmt.Println("Buah Favorite:", buahFavorite)
}
