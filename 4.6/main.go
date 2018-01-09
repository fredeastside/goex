package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "Привет,\t\t\tмир!"
	fmt.Println(str)
	fmt.Println(string(removeSpaces([]byte(str))))
}

func removeSpaces(sl []byte) []byte {
	var space rune
	out := sl[:0]
	for _, v := range sl {
		if unicode.IsSpace(rune(v)) {
			space = ' '
			continue
		}
		if space != 0 {
			out = append(out, byte(space))
			space = 0
		}
		out = append(out, v)
	}

	return out
}
