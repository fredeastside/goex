package main

import "fmt"

func main() {
	sl := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(sl)
	fmt.Println(remove(sl, 2))
	// fmt.Println(sl)
	// reverse(sl)
	// fmt.Println(sl)
	// rotateLeft(sl, 2)
	// fmt.Println(sl)
	// rotateRight(sl, 3)
	// fmt.Println(sl)
	// x := 7
	// res := 0
	// for x != 0 {
	// 	res++
	// 	fmt.Printf("%b\n", x)
	// 	fmt.Printf("%b\n", x-1)
	// 	x = x & (x - 1)
	// }
	// fmt.Println(res)
}

func remove(sl []int, n int) []int {
	copy(sl[n:], sl[n+1:])

	return sl[:len(sl)-1]
}

func reverse(sl []int) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

func rotateLeft(sl []int, n int) {
	reverse(sl[:n])
	reverse(sl[n:])
	reverse(sl)
}

func rotateRight(sl []int, n int) {
	reverse(sl)
	reverse(sl[:n])
	reverse(sl[n:])
}
