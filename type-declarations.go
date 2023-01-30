package main

import "fmt"

func main() {
	//Type Declaration/Aliases
	type NoKTP string
	type Married bool

	var noKtpIbad NoKTP = "13888299422222233"
	var marriedStatus Married = false

	fmt.Println(noKtpIbad)
	fmt.Println(marriedStatus)
}
