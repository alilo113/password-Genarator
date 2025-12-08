package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GeneratePassword() string {
	numbers := []byte("0123456789")
	lower   := []byte("abcdefghijklmnopqrstuvwxyz")
	upper   := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	symbols := []byte("!@#$%^&*()-_=+[]{}|;:,.<>?/~`")

	includeDigits := true
	includeLower  := true
	includeUpper  := false
	includeSymbols := true

	finalPool := []byte{}
	
	var length int = 20

	if includeDigits {
		finalPool = append(finalPool, numbers...)
	}
	if includeLower {
		finalPool = append(finalPool, lower...)
	}
	if includeUpper {
		finalPool = append(finalPool, upper...)
	}
	if includeSymbols {
		finalPool = append(finalPool, symbols...)
	}

	if len(finalPool) == 0 || length <= 0 {
		return ""
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		max := big.NewInt(int64(len(finalPool)))
		n, _ := rand.Int(rand.Reader, max)
		password[i] = finalPool[n.Int64()]
	}

	return string(password)
}

func main() {
	fmt.Println("Generated Password:", GeneratePassword())
}