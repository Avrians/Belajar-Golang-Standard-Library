package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println("Math Package")
	fmt.Println("============")
	fmt.Println(math.Ceil(1.49))
	fmt.Println(math.Floor(1.51))
	fmt.Println(math.Round(1.49))
	fmt.Println(math.Round(1.51))
	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}