package cmd

import (
    "crypto/rand"
    "fmt"
    "math/big"

    "github.com/spf13/cobra"
)

var length int

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate a secure password",
    Run: func(cmd *cobra.Command, args []string) {
        password := generatePassword(length)
        fmt.Println(password)
    },
}

func init() {
    rootCmd.AddCommand(generateCmd)

    // Local flag for password length
    generateCmd.Flags().IntVarP(&length, "length", "l", 16, "Length of the password")
}

// generatePassword generates a random password of length n
func generatePassword(n int) string {
    charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?"
    password := ""
    for i := 0; i < n; i++ {
        num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        password += string(charset[num.Int64()])
    }
    return password
}