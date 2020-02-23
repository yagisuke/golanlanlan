package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("0: ", num)
	} else if num == 1 {
		fmt.Println("1: ", num)
	} else {
		fmt.Println("2: ", num)
	}
}
