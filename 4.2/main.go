package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
)

func main() {
	sha := flag.Int("sha", 256, "SHA checksum of the data.")
	data := flag.String("data", "", "Data.")
	flag.Parse()
	fmt.Println(*sha, *data)
	fmt.Println(SHASum(sha, data))
}

func SHASum(sha *int, data *string) string {
	str := []byte(*data)
	if *sha == 384 {
		res := sha512.Sum384(str)

		return hex.EncodeToString(res[:])
	}
	if *sha == 512 {
		res := sha512.Sum512(str)

		return hex.EncodeToString(res[:])
	}
	res := sha256.Sum256(str)

	return hex.EncodeToString(res[:])
}
