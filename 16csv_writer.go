package main

import (
	"encoding/csv"
	"os"
)

func main16()  {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"kevin", "aldo", "bagus", "pramana"})
	_ = writer.Write([]string{"budi", "pramana", "aldo", "kevin"})
	_ = writer.Write([]string{"joko", "morgan", "aldo", "kevin"})

	writer.Flush()
}