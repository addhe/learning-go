package main

import "fmt"

func jumlahkan(angka1 int, angka2 int) int {
	return angka1 + angka2
}

func main() {
	angka1 := 5
	angka2 := 3

	total := jumlahkan(angka1, angka2)

	fmt.Println("Hasil Penjumlahan:", total)
}
