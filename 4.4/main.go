package main

import (
	"fmt"
)

func main() {
	sl := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotateLeft(sl, 2))
	fmt.Println(rotateRight(sl, 2))
}

func rotateLeft(sl []int, n int) []int {
	tmp := make([]int, n)
	copy(tmp, sl[:n])
	copy(sl[:], sl[n:])
	copy(sl[len(sl)-n:], tmp)

	return sl
}

func rotateRight(sl []int, n int) []int {
	tmp := make([]int, n)
	copy(tmp, sl[len(sl)-n:])
	copy(sl[n:], sl[:len(sl)-n])
	copy(sl[:], tmp)

	return sl
}
