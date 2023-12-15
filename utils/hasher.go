package utils

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

func Sha3Hash(input string) string {
	hash := sha3.New256()
	_, _ = hash.Write([]byte(input))
	sha3 := hash.Sum(nil)
	return fmt.Sprintf("%x", sha3)
}
