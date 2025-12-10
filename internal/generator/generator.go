package generator

import (
	"crypto/rand"
	"math/big"
)

func GeneratePassword() string {
	numbers := []byte("0123456789")
	lower   := []byte("abcdefghijklmnopqrstuvwxyz")
	upper   := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	symbols := []byte("!@#$%^&*()-_=+[]{}|;:,.<>?/~`")

	finalPool := append(numbers, lower...)
	finalPool = append(finalPool, symbols...)

	length := 20

	password := make([]byte, length)

	for i := 0; i < length; i++ {
		max := big.NewInt(int64(len(finalPool)))
		n, _ := rand.Int(rand.Reader, max)
		password[i] = finalPool[n.Int64()]
	}

	return string(password)
}