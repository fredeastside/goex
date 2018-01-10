package main

import (
	"fmt"
	"unicode"
)

func main() {
	counts := make(map[string]map[rune]int)
	invalid := 0
	//in := bufio.NewReader(os.Stdin)
	const mixed = "\b5Ὂg̀9! ℃ᾭG"
	for _, r := range mixed {
		// r, n, err := in.ReadRune()
		// if err == io.EOF {
		// 	break
		// }
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		// }
		// if r == unicode.ReplacementChar && n == 1 {
		// 	invalid++
		// 	continue
		// }
		if unicode.IsControl(r) {
			addRune(counts, "control", r)
		}
		if unicode.IsDigit(r) {
			addRune(counts, "digit", r)
		}
		if unicode.IsGraphic(r) {
			addRune(counts, "graphic", r)
		}
		if unicode.IsLetter(r) {
			addRune(counts, "letter", r)
		}
		if unicode.IsLower(r) {
			addRune(counts, "lower", r)
		}
		if unicode.IsMark(r) {
			addRune(counts, "mark", r)
		}
		if unicode.IsNumber(r) {
			addRune(counts, "number", r)
		}
		if unicode.IsPrint(r) {
			addRune(counts, "print", r)
		}
		if !unicode.IsPrint(r) {
			addRune(counts, "not_print", r)
		}
		if unicode.IsPunct(r) {
			addRune(counts, "punct", r)
		}
		if unicode.IsSpace(r) {
			addRune(counts, "space", r)
		}
		if unicode.IsSymbol(r) {
			addRune(counts, "symbol", r)
		}
		if unicode.IsTitle(r) {
			addRune(counts, "title", r)
		}
		if unicode.IsUpper(r) {
			addRune(counts, "upper", r)
		}
	}
	fmt.Print("rune\tcount\n")
	for t, tm := range counts {
		fmt.Printf("%s\n", t)
		for r, n := range tm {
			fmt.Printf("\t%q\t%d\n", r, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("Invalid: %d\n", invalid)
	}
}

func addRune(m map[string]map[rune]int, t string, r rune) {
	if m[t] == nil {
		m[t] = make(map[rune]int)
	}
	m[t][r]++
}
