package main

import (
	"fmt"
)

func main() {
	const command = "left"

	switch {
	case command == "top":
		fmt.Println("上")
	case command == "right":
		fmt.Println("右")
	case command == "bottom":
		fmt.Println("下")
	case command == "left":
		fmt.Println("左")
		fallthrough
	default:
		fmt.Println("行き止まり")
	}
}
