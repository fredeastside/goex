package main

import "fmt"

func main() {
	sl := []string{"x", "y", "y", "y", "x", "z", "z", "W", "W", "W"}
	fmt.Println(remove(sl))
}

func remove(sl []string) []string {

	out := sl[:0]
	for i, v := range sl {
		if i != len(sl)-1 && v == sl[i+1] {
			continue
		}
		out = append(out, v)
	}

	return out
}
