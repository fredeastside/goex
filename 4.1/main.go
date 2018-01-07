package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	for i, v := range c1 {
		fmt.Println(popCount(v), popCount(c2[i]))
	}
}

func popCount(x byte) int {
	return int(pc[x>>(0*8)] +
		pc[x>>(1*8)] +
		pc[x>>(2*8)] +
		pc[x>>(3*8)] +
		pc[x>>(4*8)] +
		pc[x>>(5*8)] +
		pc[x>>(6*8)] +
		pc[x>>(7*8)])
}
