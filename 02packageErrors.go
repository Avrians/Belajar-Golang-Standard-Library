package main

import "errors"

var (
	ValidationError = errors.New("Validation error")
	NotFoundError = errors.New("Not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Eko" {
		return NotFoundError
	}
	
	return nil
}

func main02()  {
	err := GetById("Eko")
	if err != nil {
		if errors.Is(err, ValidationError) {
			println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			println("Not found error")
		} else {
			println("Unknown error")
		}
	} else {
		println("Success")
	}
}