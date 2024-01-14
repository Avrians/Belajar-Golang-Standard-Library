package main

import (
	"fmt"
	"strings"
)

func main05()  {
	fmt.Println(strings.Contains("Kevin Sanjaya", "Kevin"))
	fmt.Println(strings.Split("Kevin Sanjaya", " "))
	fmt.Println(strings.ToLower("Kevin Sanjaya"))
	fmt.Println(strings.ToUpper("Kevin Sanjaya"))
	fmt.Println(strings.Trim("     Kevin Sanjaya     ", " "))
	fmt.Println(strings.ReplaceAll("Kevin Sanjaya", "Kevin", "Bagus"))
}