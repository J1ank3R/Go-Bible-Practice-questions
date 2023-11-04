package chapter4

import (
	"crypto/sha256"
	"fmt"
)

func chapter4_1() {
	var input1, input2 string
	fmt.Print("Please input the first string: ")
	fmt.Scan(&input1)
	fmt.Print("Please input the second string: ")
	fmt.Scan(&input2)

	c1 := sha256.Sum256([]byte(input1))
	c2 := sha256.Sum256([]byte(input2))

	diffBits := countDifferentBits(c1[:], c2[:])
	fmt.Println("Number of different bits", diffBits)
}

func countDifferentBits(hash1, hash2 []byte) int {
	diffBits := 0

	for i := 0; i < len(hash1) && i < len(hash2); i++ {
		diff := hash1[i] ^ hash2[i]
		for j := 0; j < 8; j++ {
			if diff&(1<<uint(j)) != 0 {
				diffBits += 1
			}
		}
	}
	return diffBits
}
