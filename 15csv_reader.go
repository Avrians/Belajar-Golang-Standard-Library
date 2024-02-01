package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main15()  {
	csvString := "kevin,aldo,bagus,pramana\n"	+
				"budi,pramana,aldo,kevin\n" +
				"joko,morgan,aldo,kevin\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}