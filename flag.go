package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")        // param 1 = name, param 2 = default value, param 3 = desc
	var user *string = flag.String("user", "root", "Put your database user")             // param 1 = name, param 2 = default value, param 3 = desc
	var password *string = flag.String("password", "root", "Put your database password") // param 1 = name, param 2 = default value, param 3 = desc
	var port *int = flag.Int("port", 3306, "Put your database port")                     // param 1 = name, param 2 = default value, param 3 = desc

	//Parsing
	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	fmt.Println("Port : ", *port)

	//Run script in terminal : go run .\flag.go -host="192.10.10.2" -user=root -password=p@$$w0rd -port=8080
}

/**
Note :
- Package flag berisikan fungsionalitas untuk memparsing command line argument
- https://golang.org/pkg/flag/
*/
