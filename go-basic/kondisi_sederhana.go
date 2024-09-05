package main

import "fmt"

func main() {
	umur := 900

	if umur >= 17 {
		fmt.Println("Kamu boleh membuat SIM karena umur kamu adalah", umur, "tahun")
	} else {
		fmt.Println("Maaf, kamu masih di bawah 17 tahun")
	}

}
