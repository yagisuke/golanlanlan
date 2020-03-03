package main

import (
	"fmt"
)

func main() {
	dwarfs := [...]string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"}

	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println("for ", i, dwarf)
	}

	for i, dwarf := range dwarfs {
		fmt.Println("range ", i, dwarf)
	}
}
