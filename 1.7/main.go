package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//body, err := ioutil.ReadAll(response.Body)
		body, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%v\n", body)
	}
}
