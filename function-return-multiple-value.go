package main

import "fmt"

func getFullname() (string, string) {
	return "Ibad", "Nurhamim"
}

func main() {
	firstName, lastName := getFullname()
	fmt.Println(firstName)
	fmt.Println(lastName)

	//Ignore
	fist, _ := getFullname()
	fmt.Println(fist)
}
