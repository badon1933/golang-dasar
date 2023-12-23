package main

import "fmt"

func main() {
	type noKTP string

	var ktpBadon noKTP = "1234567890"

	var contoh string = "1111111111"
	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(ktpBadon)
	fmt.Println(contohKTP)
}
