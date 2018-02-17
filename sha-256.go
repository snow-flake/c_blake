package main

import "golang.org/x/crypto/blake2b"
import "fmt"
import "./src"

func main() {
	output := blake2b.Sum256([]byte(""))
	fmt.Println(output)

	input := []byte("")
	src.blake256_hash(output, input, len(input))
	fmt.Println(output)
}

