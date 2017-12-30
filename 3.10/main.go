package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("123456789"))
}

func comma(s string) string {
	const sepLength = 3
	len := len(s)
	if len <= sepLength {
		return s
	}

	return comma(s[:len-sepLength]) + "," + s[len-sepLength:]
}

func commaBuff(s string) string {
	const sepLength = 3
	var buf bytes.Buffer
	rest := ""

	if dotIndex := strings.LastIndex(s, "."); dotIndex != -1 {
		rest = s[dotIndex:]
		s = s[:dotIndex]
	}

	for i, val := range s {
		if i != 0 && i%sepLength == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(val)
	}
	buf.WriteString(rest)

	return buf.String()
}
