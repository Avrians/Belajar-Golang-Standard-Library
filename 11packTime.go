package main

import (
	"fmt"
	"time"
)

func main11a()  {
	var now time.Time = time.Now()
	fmt.Println(now)
	fmt.Println(now.Local())

	var utc time.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2019-04-30 15:30:00"
	// value := "ASAL" // maka akan eror
	valueTime, err := time.Parse(formatter, value)
	if err == nil {
		fmt.Println(valueTime)
	} else {
		fmt.Println("error", err.Error())
	}


	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())

}