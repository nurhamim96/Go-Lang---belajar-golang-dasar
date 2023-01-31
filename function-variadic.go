package main

import "fmt"

func sumAll(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 90, 50, 40, 30)
	fmt.Println(total)

	//menggunakan slice sebagai parameter
	slice := []int{10, 20, 30, 40, 50}
	total = sumAll(slice...)
	fmt.Println(total)
}
