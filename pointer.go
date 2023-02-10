package main

import "fmt"

type Address struct {
	City, Province, Country string
}

//func ChangeCountryToIndonesia(address Address) {
//	address.Country = "Indonesia"
//}

// Menggunakan Pointer
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// Pass By Value di Go-Lang
	//address1 := Address{"Majalengka", "Jawa Barat", "Indonesia"}
	//address2 := address1
	//
	//address2.City = "Bandung"
	//
	//fmt.Println(address1)
	//fmt.Println(address2)

	//Dengan menggunakan pointer maka bisa menggunakan Pass By Reference di Go-Lang dengan menggunakan &
	var address1 Address = Address{"Majalengka", "Jawa Barat", "Indonesia"}
	//var address4 *Address = Address{"Majalengka", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 = new(Address)
	address4.City = "Jakarta "
	fmt.Println(address4)

	//Materi func pointer
	//var alamat = Address{
	//	City:     "Majalengka",
	//	Province: "Jawa Barat",
	//	Country:  "",
	//}
	//ChangeCountryToIndonesia(alamat)
	//fmt.Print(alamat)

	var alamat = Address{
		City:     "Majalengka",
		Province: "Jawa Barat",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

	var alamatPointer = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamatPointer)
}

/**
Note :
- Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
- Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi valuenya
- Pointer adalah kemampuan membuat reference ke lokasi data memory yang sama, tanpa menduplikasi data yang sudah ada
- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference
- Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti nama variablenya
- Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
- Semua variable yang mengacu ke data yang sama tidak akan berubah
- Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *
- Sebelumnya untuk membuat pointer dengan menggunakan operator &
- Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
- Namun function new hanya mengembalikan pointer ke data kosong artinya tidak ada data awal
- Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut.
- Oleh karena itu, jika mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
- Hal ini membuat variabel menjadi aman, karena tidak akan berubah.
- Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut.
- Untuk melakukan ini, kita juga bisa menggunakan pointer di function.
- Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator* di parameternya.
*/
