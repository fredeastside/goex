package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	algo := flag.Int("algo", 256, "SHA checksum of the data.")
	data := flag.String("data", "", "Data.")
	flag.Parse()
	fmt.Println(*algo, *data)
	fmt.Println(SHASum(algo, data))
}

func SHASum(algo *int, data *string) string {
	res := ""
	switch *algo {
	case 384:
		res = fmt.Sprintf("%x", sha512.Sum384([]byte(*data)))
	case 512:
		res = fmt.Sprintf("%x", sha512.Sum512([]byte(*data)))
	default:
		res = fmt.Sprintf("%x", sha256.Sum256([]byte(*data)))
	}

	return res
}
