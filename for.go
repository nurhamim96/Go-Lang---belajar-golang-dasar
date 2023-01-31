package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	//For dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	//Looping slice
	slice := []string{"Ibad", "Nurhamim", "Bambang", "Ucup"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//Looping slice dengan For Range
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// jika mau print valuenya saja bisa menggunakan underscore untuk variabel i yg menandakan variabel tsb tidak terpakai
	//Looping slice dengan For Range
	for _, value := range slice {
		fmt.Println(value)
	}

	//Looping map dengan For Range
	person := make(map[string]string)
	person["name"] = "Ibad"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
