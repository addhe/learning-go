package main

import "fmt"

// Fungsi Average menghitung rata-rata dari serangkaian angka.
//
// Ia menerima argumen berupa slice float64 dan mengembalikan nilai float64 tunggal
// yang merepresentasikan rata-rata dari angka-angka tersebut.
func Average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	fmt.Println(Average(xs))
}
