package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var abbr string

	phrase := readString()

	//Calculating result
	sl := strings.Fields(phrase)
	for _, val := range sl {
		if unicode.IsLetter([]rune(val)[0]) {
			abbr += string(unicode.ToUpper([]rune(val)[0]))
		}
	}

	//Writing result to Stdout
	fmt.Println(abbr)
}

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
