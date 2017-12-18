package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		// 0 = 0 0 + 0
		// 1 = 1 0 + 1
		// 1 = 2 1 + 0
		// 2 = 3 1 + 1
		// 1 = 4 1 + 0
		// 2 = 5 1 + 1
		// 2 = 6 2 + 0
		// 3 = 7 2 + 1
		// 1 = 8 1 + 0
		// 2 = 9 1 + 1
		// 2 = 10 2 + 0
		// 3 = 11 2 + 1
		// 2 = 12 2 + 0
		// 3 = 13 2 + 1
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	now := time.Now()
	fmt.Println(popCount(1023))
	fmt.Printf("%.10f", time.Since(now).Seconds())
	now = time.Now()
	fmt.Println(popCountLoop(1023))
	fmt.Printf("%.10f", time.Since(now).Seconds())
}

func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func popCountLoop(x uint64) int {
	res := byte(0)
	for i := uint64(0); i < 8; i++ {
		res += pc[byte(x>>(i*8))]
	}

	return int(res)
}
