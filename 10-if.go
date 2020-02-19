package main

import (
	"fmt"
)

func main() {
	const command = "go mountain"

	if command == "go mountain" {
		fmt.Println("今夜が山だ")
	} else if command == "go sea" {
		fmt.Println("明日は海だ")
	} else {
		fmt.Println("週末も会社だ")
	}
}
