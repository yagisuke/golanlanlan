package main

import (
	"fmt"
	"reflect"
)

func main() {
	days1 := 365.2425
	fmt.Println(reflect.TypeOf(days1), days1)

	var days2 = 365.2425
	fmt.Println(reflect.TypeOf(days2), days2)

	var days3 float64 = 365.2425
	fmt.Println(reflect.TypeOf(days3), days3)

	var days4 float64 = 365
	fmt.Println(reflect.TypeOf(days4), days4)

	var days5 float64 = 365.0
	fmt.Println(reflect.TypeOf(days5), days5)
}
