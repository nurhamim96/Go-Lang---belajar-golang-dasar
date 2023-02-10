package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Ibad")
	//helper.sayGoodbye("Ibad") //error
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) //error
}

/**
Note :
- GOPATH -> environment variable yang berisikan lokasi tempat kita menyimpan project dan library Go-Lang
- GOPATH wajib di buat ketika kita mulai membuat aplikasi lebih dari satu file atau butuh menggunakan library
- Package -> tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di Go-Lang
- Dengan menggunakan package, kita bisa merapikan kode program yang kita buat
- Package sendiri sebenarnya hanya direktori folder di sistem operasi kita
- Secara standart, file Go-Lang hanya bisa mengakses file Go-Lang yang berada dalam package yang sama
- Jika kita ingin mengakses file Go-Lang yang berada diluar package, maka kita bisa menggunakan import
- Dibahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable
- Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable
- Jika nama nya diawali dengan huruf besar, maka artinya bisa diakses dari package lain, jika dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain
*/
