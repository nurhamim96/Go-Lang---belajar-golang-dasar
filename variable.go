package main

import "fmt"

func main() {
	var name string

	name = "Ibad Nur"
	fmt.Println(name)

	name = "Ibad Nurhamim"
	fmt.Println(name)

	var friendName = "Ucup"
	fmt.Println(friendName)

	var age = 20
	fmt.Println(age)

	//Tanpa var (digunakan hanya untuk deklarasi awal)
	country := "Indonesia"
	fmt.Println(country)

	//Deklarasi Multiple Variable
	var (
		firstName = "Ibad"
		lastName  = "Nurhamim"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
