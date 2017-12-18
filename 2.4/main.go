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
	var count int
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}

	return count
}
