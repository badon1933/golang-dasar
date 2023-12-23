package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b
	var d = c - 1
	var e = a + b - a/2 + a*b
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	/* Augmented assignment */
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)
	i += 5 // i = i + 5
	fmt.Println(i)

	/* Unary Operator */
	var j = 0
	j++ // j = j + 1
	fmt.Println(j)
	j++ // j = j + 1
	fmt.Println(j)

	j-- // j = j - 1
	fmt.Println(j)
}
