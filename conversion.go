package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	// hati" jika konversi ke yang lebih redah, karena akan merubah value jika tipe data tidak sanggup menghandle
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// Konversi dari byte ke string
	var name = "Ibad"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(eString)

}
