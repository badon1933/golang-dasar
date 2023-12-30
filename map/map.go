package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Badon"
	// person["age"] = "25"
	// person["address"] = "Purwakarta"

	/* Deklarasi cara cepat */

	person := map[string]string{
		"name":    "Badon",
		"age":     "25",
		"address": "Purwakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])

	/* Operasi Map */
	fmt.Println("Panjang Map person adalah", len(person))
	delete(person, "address")
	fmt.Println(person)
}
