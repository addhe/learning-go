package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan angka pertama dan kedua
	var angka1, angka2 float64

	// Meminta angka pertama dan kedua dari pengguna
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angka1)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)

	// Mencetak hasil penjumlahan
	fmt.Printf("Hasil penjumlahan: %.2f + %.2f = %.2f\n", angka1, angka2, angka1+angka2)

	// Mencetak hasil pengurangan
	fmt.Printf("Hasil pengurangan: %.2f - %.2f = %.2f\n", angka1, angka2, angka1-angka2)

	// Mencetak hasil perkalian
	fmt.Printf("Hasil perkalian: %.2f * %.2f = %.2f\n", angka1, angka2, angka1*angka2)

	// Mencetak hasil pembagian (dengan penanganan error jika angka kedua adalah 0)
	if angka2 != 0 {
		fmt.Printf("Hasil pembagian: %.2f / %.2f = %.2f\n", angka1, angka2, angka1/angka2)
	} else {
		fmt.Println("Error: tidak dapat membagi dengan 0")
	}
}
