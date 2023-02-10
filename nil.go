package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := newMap("Ibad")
	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person)
	}
}

/**
Note:
- Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nill
- Berbeda dengan Go-Lang saat kita buat variabel dengan tipe data tertentu, maka secara otomatis akan dibuatkan default valuenya
- Namun di Go-Lang ada data nil, yaitu data kosong
- Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel
*/
