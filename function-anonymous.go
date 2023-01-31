package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// Menggunakan func biasa
//func blackListAdmin(name string) bool {
//	return name == "admin"
//}
//
//func blackListRoot(name string) bool {
//	return name == "root"
//}

func main() {
	//Menggunakan anonymous function
	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blackList)
	registerUser("ibad", blackList)

	//Menggunakan anonymous function pada parameter
	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("ibad", func(name string) bool {
		return name == "root"
	})
}
