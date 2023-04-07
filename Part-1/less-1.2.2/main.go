package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64

	// Reading parameters from Stdin
	_, err := fmt.Scan(&x1, &y1, &x2, &y2)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Calculating result
	d := math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))

	//Writing result to Stdout
	fmt.Println(d)
}
