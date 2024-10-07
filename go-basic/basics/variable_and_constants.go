package main

import "fmt"

func main() {
	// Deklarasi variabel integer
	var angkaInteger int = 10
	fmt.Println("Nilai angka integer:", angkaInteger)

	// Deklarasi variabel float
	var angkaFloat float64 = 3.14
	fmt.Println("Nilai angka float:", angkaFloat)

	// Deklarasi variabel string
	var namaString string = "John Doe"
	fmt.Println("Nilai nama string:", namaString)

	// Deklarasi variabel boolean
	var nilaiBoolean bool = true
	fmt.Println("Nilai nilai boolean:", nilaiBoolean)

	// Deklarasi variabel tanpa jenis (Go akan menentukan jenisnya secara otomatis)
	var angkaTanpaJenis = 20
	fmt.Println("Nilai angka tanpa jenis:", angkaTanpaJenis)

	// Deklarasi variabel dengan sintaks pendek (:=)
	angkaPendek := 30
	fmt.Println("Nilai angka pendek:", angkaPendek)
}
