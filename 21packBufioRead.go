package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main21()  {
	input := strings.NewReader("Hello, this is long string\nthat is split over\nmany lines\n")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
