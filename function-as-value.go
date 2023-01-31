package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	sayGoogBye := getGoodBye
	result := sayGoogBye("Ibad")
	fmt.Println(result)
	fmt.Println(getGoodBye("Ibad"))
}
