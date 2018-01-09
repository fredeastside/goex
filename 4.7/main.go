package main

import (
	"fmt"
)

func main() {
	str := "Привет, мир!"
	fmt.Println(str)
	fmt.Println(string(reverse([]byte(str))))
}

func reverse(sl []byte) []byte {
	// prev := 0
	// for i := range string(sl) {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(string(sl[prev:i]))
	// 	prev = i
	// }

	// fmt.Println(string(sl[prev:len(sl)]))
	runes := []rune(string(sl))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}
