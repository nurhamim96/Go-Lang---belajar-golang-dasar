package main

import "fmt"

func main() {
	var name = "Ibad"

	if name == "Ibad" {
		fmt.Println("Hello Ibad")
	} else if name == "Ucup" {
		fmt.Println("Hello Ucup")
	} else if name == "Bambang" {
		fmt.Println("Hello Bambang")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	//If dengan short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
