package commands

import (
    "crypto/rand"
    "fmt"
    "math/big"
)

// Accept length as argument
func runGenerate(length int) string {
	password := generatePassword(length)
	fmt.Println("Generated password:", password)
	return password
}

func generatePassword(n int) string {
    charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?"
    password := ""
    for i := 0; i < n; i++ {
        num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        password += string(charset[num.Int64()])
    }
    return password
}