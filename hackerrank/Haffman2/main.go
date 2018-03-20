package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// f, err := os.Open("input.txt")
	// if err != nil {
	// 	fmt.Errorf("Error: %v", err)
	// }
	// defer f.Close()
	scanner := bufio.NewScanner(os.Stdin)

	// var lettersCount string
	// var codeLength string
	var code string
	m := make(map[string]string)

	for scanner.Scan() {
		txt := scanner.Text()
		lineSl := strings.Split(txt, ":")
		if len(lineSl) == 2 {
			m[strings.TrimSpace(lineSl[1])] = lineSl[0]
			continue
		}

		lineSl = strings.Split(txt, " ")
		if len(lineSl) == 2 {
			// lettersCount = lineSl[0]
			// codeLength = lineSl[1]
			continue
		}

		code = txt
	}

	var letterCode string
	for len(code) != 0 {
		for c, letter := range m {
			if !strings.HasPrefix(code, c) {
				continue
			}
			letterCode += letter
			runes := []rune(code)
			code = string(runes[len(c):])
		}
	}
	fmt.Println(letterCode)
}
