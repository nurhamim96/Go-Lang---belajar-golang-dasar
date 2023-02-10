package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("i([a-z])([a-z])d")

	fmt.Println(regex.MatchString("ibad"))
	fmt.Println(regex.MatchString("ipad"))
	fmt.Println(regex.MatchString("iPad"))

	search1 := regex.FindAllString("ibad ipad  idad", 2)
	fmt.Println(search1)

	search2 := regex.FindAllString("ibad ipad  idad", -1)
	fmt.Println(search2)
}

/**
Note :
- Package regexp -> utilitas di Go-Lang untuk melakukan pencarian regular expression
- Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
- https://github.com/google/re2/wiki/Syntax
- https://golang.org/pkg/regexp/
*/
