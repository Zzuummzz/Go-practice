package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	r := make([]int, 0, len(iterable))
	for _, val := range iterable {
		if predicate(val) {
			r = append(r, val)
		}
	}
	return r
}

func fEven(d int) bool {
	return d%2 == 0
}

func main() {
	// Reading parameters from Stdin
	src := readInput()

	//Calculating result
	res := filter(fEven, src)

	//Writing result to Stdout
	fmt.Println(res)
}

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
