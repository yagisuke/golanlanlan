package main

import (
	"fmt"
)

func main() {
	const age = 31
	const young = age < 20
	fmt.Println("未成年ですか？", young)
}
