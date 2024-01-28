package main

import (
	"fmt"
	"regexp"
)

func main()  {
	var regex *regexp.Regexp = regexp.MustCompilePOSIX(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	fmt.Println(regex.FindAllString("eko eka edo eki e1o", 10))
}