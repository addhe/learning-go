package main

import "fmt"

func main() {
	nama_hari := [7]string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}

	//fmt.Println("Hari ke-2:", nama_hari[1])

	for i, hari := range nama_hari {
		fmt.Println("Hari ke -", i+1, ":", hari)
	}
}
