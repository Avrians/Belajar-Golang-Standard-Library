package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string `required:"true" max:"10"`
	Address string `required:"true"`
	Email string `required:"true"`
}

func readField(value any)  {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type Name: ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main12()  {
	readField(Sample{Name: "Sample"})
	readField(Person{"Joko", "jl", "joko@gmail.com"})

	person := Person{
		Name: "Joko",
		Address: "jl",
		Email: "kkk",
	}
	fmt.Println(isValid(person))
}