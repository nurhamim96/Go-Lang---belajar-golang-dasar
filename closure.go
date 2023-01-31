package main

import "fmt"

func main() {
	name := "Bambang"
	counter := 0

	increment := func() {
		name = "Ucup"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
