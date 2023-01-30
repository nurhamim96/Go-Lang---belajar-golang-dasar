package main

import "fmt"

func main() {
	name := "Ibad"

	switch name {
	case "Ibad":
		fmt.Println("Hallo Ibad")
		fmt.Println("Hallo Ibad")
	case "Bambang":
		fmt.Println("Hallo Bambang")
		fmt.Println("Hallo Bambang")
	case "Ucup":
		fmt.Println("Hallo Ucup")
		fmt.Println("Hallo Ucup")
	default:
		fmt.Println("Hi, kenalan donk")
		fmt.Println("Hi, kenalan donk")
	}

	//Switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	//Switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
