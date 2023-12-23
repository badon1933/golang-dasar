package main

import "fmt"

func main() {
	var nilai32 int32 = 32770
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Badon Ganteng"
	var b uint8 = name[0]
	var bString = string(b)

	fmt.Println(name)
	fmt.Println(b)
	fmt.Println(bString)
}
