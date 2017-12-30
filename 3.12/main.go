package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(anagram("123", "321"))
}

func anagram(s1 string, s2 string) bool {
	sl1 := strings.Split(s1, "")
	sl2 := strings.Split(s2, "")
	sort.Strings(sl1)
	sort.Strings(sl2)

	return strings.Join(sl1, "") == strings.Join(sl2, "")
}
