package main

import (
	"fmt"
)

func main() {
	celsius1 := 21.0
	fmt.Println((celsius1/5.0*9.0)+32, "°F")
	fmt.Println((9.0/5.0*celsius1)+32, "°F")

	celsius2 := 21.0
	fmt.Println((celsius2*9.0/5.0)+32, "°F")
	fmt.Println((9.0*celsius2/5.0)+32, "°F")
}
