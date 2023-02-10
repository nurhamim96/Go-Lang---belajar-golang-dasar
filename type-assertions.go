package main

import "fmt"

func random() interface{} {
	return "Ibad"
}

func main() {
	var result interface{} = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}

/**
Note:
- Type assertions merupakan kemampuan tipe data menjadi tipe data yang diinginkan
- Fitur ini sering kali digunakan ketika kita bertemu dengan data interface kosong
- Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
- Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
- Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
*/
