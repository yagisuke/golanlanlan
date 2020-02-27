package main

import (
	"fmt"
)

func main() {
	var red, green, blue uint8 = 0x00, 0x8d, 0xd5
	fmt.Printf("%x %x %x\n", red, green, blue)

	red, green, blue = 0, 141, 213
	fmt.Printf("%x %x %x\n", red, green, blue)

	fmt.Printf("color: #%02x%02x%02x", red, green, blue)
}
