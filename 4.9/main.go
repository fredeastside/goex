package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Print("word:\tcount:\n")
	for w, c := range wordfreq(f) {
		fmt.Printf("%s\t%d\n", w, c)
	}
}

func wordfreq(f *os.File) map[string]int {
	words := make(map[string]int)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words[scanner.Text()]++
	}

	return words
}
