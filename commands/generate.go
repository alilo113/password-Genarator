package commands

import (
    "crypto/rand"
    "fmt"
    "math/big"
)

// Accept length as argument
func runGenerate(length int) {
    password := generatePassword(length)
    fmt.Println("Generated Password:", password)
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