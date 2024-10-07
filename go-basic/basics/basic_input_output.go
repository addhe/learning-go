package main

import "fmt"

func main() {
	// Deklarasi variable untuk menyimpan nama pengguna
	var name string

	// Meminta nama pengguna
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&name)

	// Menyapa pengguna dengan pesan yang di personalisasi
	fmt.Printf("Halo, %s! Selamat datang!\n", name)

}
