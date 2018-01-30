package main

import (
	"flag"
	"fmt"

	"github.com/fredeastside/goex/4.12/xkcd"
)

func main() {
	var number int
	flag.IntVar(&number, "num", 0, "Number of comics.")
	flag.Parse()
	comic, err := xkcd.FindByNumber(number)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Comic text: %s", comic.Transcript)
}
