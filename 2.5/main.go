package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	now := time.Now()
	fmt.Println(popCount(1023))
	fmt.Printf("%.10f", time.Since(now).Seconds())
}

func popCount(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
