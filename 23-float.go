package main

import (
	"fmt"
)

func main() {
	third := 1.0 / 3.0
	result := third + third + third
	fmt.Printf("%T型 %[1]v\n", result)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Printf("%T型 %[1]v\n", piggyBank)
}
