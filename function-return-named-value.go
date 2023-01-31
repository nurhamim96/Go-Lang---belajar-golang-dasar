package main

import "fmt"

func getFullName2() (firstName string, lastName string, nickName string) {
	firstName = "Ibad"
	lastName = "Nurhamim"
	nickName = "IbadNur"
	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
