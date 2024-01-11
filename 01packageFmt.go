package main

import "fmt"


func main() {
	fmt.Println("Hello, World!")

	firstName := "John"
	lastName := "Doe"
	fmt.Println("Hallo '",firstName, lastName, "'")
	fmt.Printf("Hallo '%s %s'\n", firstName, lastName)
}