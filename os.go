package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	// run in terminal : go run  .\os.go ibad nurhamim ibadnur

	//Get Hostname
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	gopath := os.Getenv("GOPATH")
	fmt.Println(gopath)
}

/**
Note :
- Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
- Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan disemua sistem operasi)
- https://golang.org/pkg/os/
*/
