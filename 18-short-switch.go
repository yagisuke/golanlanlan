package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch num := rand.Intn(3); num {
	case 0:
		fmt.Println("0: ", num)
	case 1:
		fmt.Println("1: ", num)
	case 2:
		fmt.Println("2: ", num)
	default:
		fmt.Println("error: error.")
	}
}
