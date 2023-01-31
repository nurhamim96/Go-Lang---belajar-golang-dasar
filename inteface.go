package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHi(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var ibad Person
	ibad.Name = "Ibad"

	sayHi(ibad)

	cat := Animal{
		Name: "Push",
	}

	sayHi(cat)
}

/**
Note :
- Interface -> tipe data Abstract, dia tidak memiliki implementasi langsung
- Sebuah interface berisikan definisi-definisi langsung
- Biasanya interface digunakan sebagai kontrak
- Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
- Sehingga kita tidak perlu mengimplementasikan interface secara manual
- Hal ini agak berbeda dengan bahasa pemrograman lain ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface yang mana
*/
