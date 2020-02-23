package main

import (
	"fmt"
)

func main() {
	var a = 0
	for a = 10; a > 0; a-- {
		fmt.Println("a: ", a)
	}

	for b := 10; b > 0; b-- {
		fmt.Println("b: ", b)
	}
}
