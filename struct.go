package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuu from", a.Name)
}

func main() {
	var data Customer
	data.Name = "Ibad Nurhamim"
	data.Address = "Jl. Nin aja dulu"
	data.Age = 20

	data.sayHi("Bambang")
	data.sayHuuu()

	//fmt.Println(data.Name)
	//fmt.Println(data.Address)
	//fmt.Println(data.Age)
	//
	//ucup := Customer{
	//	Name:    "Ucup",
	//	Address: "Jl. Kaki",
	//	Age:     35,
	//}
	//
	//fmt.Println(ucup)
	//
	//bambang := Customer{"Bambang", "Cirebon", 35}
	//fmt.Println(bambang)
}

/**
Note :
- Struct -> sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
- Struct biasanya representasi data dalam program aplikasi yang kita buat
- Data di struct disimpan dalam field
- Sederhananya struct adalah kumpulan dari field
- Struct tidak bisa langsung digunakan
- Namun kita bisa membuat data/object dari struct yang telah kita buat
- Struct adalah tipe data seperti data lainnya, dia bisa digunakan sebagai parameter untuk func
*/
