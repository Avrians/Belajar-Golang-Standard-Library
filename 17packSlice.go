package main

import (
	"fmt"
	"slices"
)


func main17()  {
	names := []string{"kevin", "aldo", "bagus", "pramana"}
	values := []int{100, 90, 30, 70}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "kevin"))
	fmt.Println(slices.Contains(names, "joko"))
	fmt.Println(slices.Contains(values, 100))
	fmt.Println(slices.Contains(values, 50))
	fmt.Println(slices.Index(names, "kevin"))
	fmt.Println(slices.Index(names, "joko"))
}