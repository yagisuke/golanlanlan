package main

import (
	"fmt"
)

func main() {
	var count int = 30

	age := func() int {
		return count
	}

	fmt.Println(age())

	count++

	fmt.Println(age())
}
