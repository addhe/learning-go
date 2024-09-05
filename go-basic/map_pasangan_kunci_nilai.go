package main

import "fmt"

func main() {
	nilaiSiswa := map[string]int{
		"Alice":   85,
		"Bob":     92,
		"Charlie": 78,
	}

	// Mengakses nilai menggunakan kunci
	fmt.Println("Nilai Alice:", nilaiSiswa["Alice"])

	fmt.Println("Nilai Bob:", nilaiSiswa["Bob"])
	fmt.Println("Nilai Charlie:", nilaiSiswa["Charlie"])

	// Menambahkan pasangan kunci-nilai baru
	nilaiSiswa["Eve"] = 88

	fmt.Println("Nilai Eve:", nilaiSiswa["Eve"])

	// Menghapus pasangan kunci-nilai
	delete(nilaiSiswa, "Alice")

	// Iterasi melalui map menggunakan for range loop
	fmt.Println("Semua Nilai Siswa:")
	for nama, nilai := range nilaiSiswa {
		fmt.Println(nama, ":", nilai)
	}
}
