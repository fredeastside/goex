package main

import (
	"fmt"
	"os"
)

func main() {
	args := ""
	for i := 1; i < len(os.Args); i++ {
		args += os.Args[i]
		args += " "
	}
	fmt.Println("Name: ", os.Args[0])
	fmt.Println("Args: ", args)
}
