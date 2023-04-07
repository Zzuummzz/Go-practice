package main

import (
	"fmt"
)

func main() {
	var code string

	// Reading parameters from Stdin
	_, err := fmt.Scan(&code)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Calculating result
	switch code {
	case "en":
		code = "English"
	case "fr":
		code = "French"
	case "ru", "rus":
		code = "Russian"
	default:
		code = "Unknown"
	}

	//Writing result to Stdout
	fmt.Println(code)
}
