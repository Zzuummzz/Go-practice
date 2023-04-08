package main

import (
	"fmt"
)

func main() {
	var (
		text, res string
		width     int
	)

	// Reading parameters from Stdin
	_, err := fmt.Scanf("%s %d", &text, &width)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Calculating result
	sl := []byte(text)
	switch {
	case width >= len(sl):
		res = text
	case width < len(sl):
		res = string(sl[:width]) + "..."

	}

	//Writing result to Stdout
	fmt.Println(res)
}
