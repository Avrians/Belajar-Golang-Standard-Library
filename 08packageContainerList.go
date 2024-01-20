package main

import (
	"container/list"
	"fmt"
)

func main()  {
	var data *list.List = list.New()

	data.PushBack("Aku")
	data.PushBack("Kamu")
	data.PushBack("Dia")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Aku

	next := head.Next() // Kamu
	fmt.Println(next.Value)

	next = next.Next() // Dia
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}