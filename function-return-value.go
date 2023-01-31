package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello, Salam kenal"
	} else {
		return "Hello, " + name
	}
}

func main() {
	result := getHello("Ibad")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
