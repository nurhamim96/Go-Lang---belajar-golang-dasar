package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "****"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Ibad", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	//menggunakan func as value
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
