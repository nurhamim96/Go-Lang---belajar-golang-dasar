package main

import "fmt"

func main() {
	//const firstName string = "Ibad"
	//const lastName = "Nurhamim"
	//const value = 1000

	//Deklarasi Multiple Constant
	const (
		firstName string = "Ibad"
		lastName         = "Nurhamim"
		value            = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}
