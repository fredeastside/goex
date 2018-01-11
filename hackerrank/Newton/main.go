package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	for {
		res := float64(z - float64((math.Pow(z, 2)-x)/float64(2*z)))
		if z-res < 0.00001 {
			break
		}
		z = res
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
