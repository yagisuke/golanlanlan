package main

import "fmt"

func main() {
	const (
		distance = 56000000
		speed = 100800
	)
	fmt.Println(distance / speed, "sec")
}
