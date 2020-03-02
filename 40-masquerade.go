package main

import (
	"fmt"
)

var f1 = func() {
	fmt.Println("Hello World. 1")
}

func main() {
	f1()

	f2 := func(message string) {
		fmt.Println(message)
	}
	f2("Hello World. 2")

	func() {
		fmt.Println("Hello World. 3")
	}()
}
