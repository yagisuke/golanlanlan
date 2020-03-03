package main

import (
	"fmt"
)

func main() {
	days := [...]string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"}

	days2 := days

	days[2] = "whoops"

	fmt.Println("days: ", days)
	fmt.Println("days2: ", days2)
}
