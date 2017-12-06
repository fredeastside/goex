package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(os.Stdout, response.Body)
	response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("fetch чтение %s: %v\n", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs elapsed\n %7d %s", time.Since(start).Seconds(), nbytes, url)
}
