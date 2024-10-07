package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan nilai kondisi
	var kondisi1, kondisi2 bool

	// Mengatur nilai kondisi
	kondisi1 = true
	kondisi2 = true

	// Memeriksa apakah kedua kondisi adalah benar
	if kondisi1 && kondisi2 {
		fmt.Println("Kedua kondisi adalah benar")
	} else {
		fmt.Println("Kedua kondisi tidak benar")
	}
}
