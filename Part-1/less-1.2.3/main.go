package main

import (
	"fmt"
)

func main() {
	var (
		source, result string
		times          int
	)

	// Reading parameters from Stdin
	_, err := fmt.Scan(&source, &times)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Calculating result
	for i := 0; i < times; i++ {
		result += source
	}

	//Writing result to Stdout
	fmt.Println(result)
}
