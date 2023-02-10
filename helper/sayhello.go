package helper

import "fmt"

var version = 1 //Tidak bisa di akses dari luar package
var Application = "Belajar Go-Lang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// Tidak bisa di akses dari luar package
func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
