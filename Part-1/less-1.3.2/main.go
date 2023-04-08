package main

import (
	"fmt"
	"sort"
)

func main() {
	var number int

	// Reading parameters from Stdin
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Calculating result
	counter := make(map[int]int)
	for ; number > 0; number /= 10 {
		counter[number%10]++
	}

	//Writing result to Stdout
	printCounter(counter)
}

// printCounter печатает карту в формате
// key1:val1 key2:val2 ... keyN:valN
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
