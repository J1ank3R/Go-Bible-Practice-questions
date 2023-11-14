package chapter4

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func chapter4_2() {
	methods := flag.String("methods", "sha256", "Hash methods (sha256, sha384, sha512)")
	flag.Parse()
	println("methods: ", *methods)
	var input string
	fmt.Scan(&input)

	switch *methods {
	case "sha256":
		hashResult := sha256.Sum256([]byte(input))
		fmt.Printf("Sha256: %x\n", hashResult)
	case "sha384":
		hashResult := sha512.Sum384([]byte(input))
		fmt.Printf("sha384: %x\n", hashResult)
	case "sha512":
		hashResult := sha512.Sum512([]byte(input))
		fmt.Printf("sha512: %x\n", hashResult)
	default:
		fmt.Println("unknow hash methods")
		os.Exit(1)
	}
}
