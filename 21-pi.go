package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var pi64 float64 = math.Pi
	fmt.Println(reflect.TypeOf(pi64), pi64)

	var pi32 float32 = math.Pi
	fmt.Println(reflect.TypeOf(pi32), pi32)
}
