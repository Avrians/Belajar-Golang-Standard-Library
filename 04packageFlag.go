package main

import (
	"flag"
	"fmt"
)

func main04()  {
	var username *string = flag.String("username", "root", "database username for mysql")
	var password *string = flag.String("password", "root", "database password for mysql")
	var host *string = flag.String("host", "localhost", "database host for mysql")
	var port *int = flag.Int("port", 3306, "database port for mysql")

	flag.Parse()

	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
	fmt.Println("Host :", *host)
	fmt.Println("Port :", *port)

}