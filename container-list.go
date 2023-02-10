package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Ibad")     //menambahkan data ke paling belakang
	data.PushBack("Nurhamim") //menambahkan data ke paling belakang
	data.PushBack("IbadNur")  //menambahkan data ke paling belakang
	data.PushFront("Bambang") //menambahkan data ke paling depan

	//ambil data selanjutnya
	fmt.Println(data.Front().Next().Value)

	//ambil data sebelumnya
	fmt.Println(data.Back().Prev().Value)

	//get data paling depan
	fmt.Println(data.Front().Value)
	//get data paling belakang
	fmt.Println(data.Back().Value)

	//looping double linked-list
	//fmt.Println("----------------- Loop -----------------")
	//for element := data.Front(); element != nil; element = element.Next() {
	//	fmt.Println(element.Value)
	//}

	//Reverse loop
	fmt.Println("----------------- Reverse Loop -----------------")
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}

/**
Note :
- Package container/list adalah implementasi struktur data double linked list di Go-Lang
- https://golang.org/pkg/container/list/
*/
