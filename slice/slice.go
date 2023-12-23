package main

import "fmt"

func main() {
	names := [...]string{
		"Doni", "Prasetya", "Arman", "Wahyu", "Budi", "Waluyo",
	}
	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	/* Operasi Slice */
	days := [...]string{
		"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu",
	}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Dinten Sabtu"
	daySlice1[1] = "Dinten Minggon"
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Hari Sabtu"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	/* Deklarasi Slice menggunakan Make */
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Slice"
	newSlice[1] = "Baru"
	//newSlice[2] = "Banget" // Error, harus menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Badon")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Endog"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	/* Copy Slice */
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	/* Hati-hari saay membuat Array */
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
