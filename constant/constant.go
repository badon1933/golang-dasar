package main

import "fmt"

func main() {
	// Meskipun tidak digunakan, tidak akan error. Tapi hanya muncul warning saja.
	const firstName string = "Badon"
	const lastName = "Ganteng"

	/* Error */
	// firstName = "Ganti"
	// lastName = "Ganti lagi"

	const (
		namaDepan    = "Badon"
		namaBelakang = "Ganteng"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)

}
