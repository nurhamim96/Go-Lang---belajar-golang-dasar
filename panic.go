package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}

//Note :
//Panic func -> function yang bisa digunakan untuk menghentikan program
//Panic func biasanya dipanggil ketika terjadi error pada saat program kita berjalan
//saat panic func dipanggil, program akan terhenti,namun defer func akan tatap di eksekusi
//Recover -> function yang bisa digunakan untuk menangkap data panic
//Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
