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

	//Print duration in minutes
	fmt.Println(s, " = ", d.Minutes())
}
