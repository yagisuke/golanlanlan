package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	var year = t.Year()
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println(year, "は閏年です")
	} else {
		fmt.Println(year, "は平年です")
	}
}
