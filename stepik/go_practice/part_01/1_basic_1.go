package main

import (
	"fmt"
	"time"
)

func main() {

	//1h30m = 90 min
	//300s = 5 min
	//10m = 10 min

	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)

	fmt.Printf("%v = %v min\n", s, d.Minutes())
}
