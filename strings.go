package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ibad Nurhamim", "Ibad"))
	fmt.Println(strings.Contains("Ibad Nurhamim", "Bambang"))

	fmt.Println(strings.Split("Ibad Nurhamim", " "))

	fmt.Println(strings.ToLower("Ibad Nurhamim"))
	fmt.Println(strings.ToUpper("Ibad Nurhamim"))
	fmt.Println(strings.ToTitle("Ibad Nurhamim"))

	fmt.Println(strings.Trim("    Ibad Nurhamim   ", " "))
	fmt.Println(strings.ReplaceAll("Bambang Bambang Bambang", "Bambang", "Ucup"))
}

/**
Note :
- Package strings -> package yang berisikan function-function untuk memanipulasi tipe data String
- https://golang.org/pkg/strings/
*/
