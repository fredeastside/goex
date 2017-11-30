package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	argsByLoop()
	fmt.Println(time.Since(start).Nanoseconds())
	start = time.Now()
	argsByJoin()
	fmt.Println(time.Since(start).Nanoseconds())
}

func argsByLoop() {
	args := ""
	for i := 1; i < len(os.Args); i++ {
		args += os.Args[i]
		args += " "
	}
	fmt.Println(args)
}

func argsByJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
