package main

import (
	"fmt"
	"time"
)

func main() {
	// Reading time duration from Stdin
	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)

	//Print splitted duration
	fmt.Println(s, " = ", d.Minutes())
}
