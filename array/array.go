package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Badon"
	names[1] = "Ganteng"
	names[2] = "Banget"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	/* Deklarasi langsung */
	var values = [3]int{
		11, 12, 13,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	/* Operasi Array */
	fmt.Println("Panjang array names =", len(names))
	fmt.Println(values)
}
