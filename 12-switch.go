package main

import (
	"fmt"
)

func main() {
	const command = "right"

	switch command {
	case "top":
		fmt.Println("上")
	case "right":
		fmt.Println("右")
	case "bottom":
		fmt.Println("下")
	case "left":
	default:
		fmt.Println("左")
	}
}
