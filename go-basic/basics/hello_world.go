// Ini adalah deklarasi paket. Di Go, setiap program harus milik sebuah paket.
// Paket "main" adalah spesial, karena merupakan titik awal program.
package main

// Baris ini mengimpor paket "fmt", yang menyediakan fungsi untuk operasi I/O yang diformat.
// Paket "fmt" adalah bagian dari perpustakaan standar Go, sehingga Anda tidak perlu menginstal apa pun tambahan untuk menggunakannya.
import "fmt"

// Ini adalah fungsi main, yang merupakan titik awal program.
// Fungsi "main" dipanggil ketika program dimulai, dan merupakan tempat di mana eksekusi program dimulai.
func main() {
	// Baris ini menggunakan fungsi "Println" dari paket "fmt" untuk mencetak "Hello, World!" ke konsol.
	// Fungsi "Println" mirip dengan "Print", tetapi juga menambahkan karakter baris baru di akhir output.
	fmt.Println("Hello, World!")
}
