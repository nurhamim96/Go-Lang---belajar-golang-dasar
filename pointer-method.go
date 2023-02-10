package main

import "fmt"

type Man struct {
	Name string
}

//func (man Man) Married() {
//	man.Name = "Mr . " + man.Name
//}

// Pointer Method
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	ibad := Man{"Ibad"}
	ibad.Married()

	fmt.Println(ibad.Name)
}

/**
Note :
- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
*/
