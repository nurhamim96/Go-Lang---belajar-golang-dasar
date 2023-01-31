package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func main() {
	sayHello()
	sayHello()
	sayHello()

	//Loop
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
