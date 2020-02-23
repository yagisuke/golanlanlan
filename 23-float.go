package main

import (
	"fmt"
	"reflect"
)

func main() {
	third := 1.0 / 3.0
	result := third + third + third
	fmt.Println(reflect.TypeOf(result), result)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)
}
