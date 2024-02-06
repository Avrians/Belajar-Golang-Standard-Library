package main

import (
	"bufio"
	"os"
)

func main22()  {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello, dunia\n")
	_, _ = writer.WriteString("Hello, golang\n")
	writer.Flush()

}