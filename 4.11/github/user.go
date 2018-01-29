package github

import "fmt"

func User() {
	response, err := Get("user")
	fmt.Println(response)
	fmt.Println(err)
}
