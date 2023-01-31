package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(10)
}

//Note :
//defer func -> digunakan untuk memanggil func setelah func lain selesai di eksekusi
//defer func akan selalu dieksekusi walaupun terjadi error di func yang di eksekusi
