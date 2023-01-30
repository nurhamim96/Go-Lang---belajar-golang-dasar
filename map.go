package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Ibad",
		"address": "Jl. Nin aja dulu",
	}

	//Menambahkan data di Map
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//Membuat Map tanpa data
	book := make(map[string]string) //var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Ibad"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
