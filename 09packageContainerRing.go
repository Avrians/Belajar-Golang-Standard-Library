package main

import (
	"container/ring"
	"fmt"
)

func main09()  {
	var data *ring.Ring = ring.New(5)

	// Manual

	// data.Value = "Value 1"

	// data = data.Next()
	// data.Value = "Value 2"

	// data = data.Next()
	// data.Value = "Value 3"

	// data = data.Next()
	// data.Value = "Value 4"

	// data = data.Next()
	// data.Value = "Value 5"

	// Automatic
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + fmt.Sprint(i)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}