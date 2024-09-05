package main

import "fmt"

type Mahasiswa struct {
	Nama string
	NIM  int
	IPK  float64
}

func main() {
	mahasiswa1 := Mahasiswa{
		Nama: "Awan",
		NIM:  12345678,
		IPK:  3.89,
	}

	fmt.Println("Nama:", mahasiswa1.Nama)
	fmt.Println("IPK:", mahasiswa1.IPK)
}
