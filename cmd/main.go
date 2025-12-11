package main

import (
	"fmt"
	"pwman/internal/entropy"
	"pwman/internal/generator"
)

func main() {
	// Generate a password of length 20
	password := generator.GeneratePassword(20)

	// Calculate pool size
	poolSize := entropy.CharSizeSet(password)

	// Calculate total entropy
	total := entropy.TotalEntropy(password, poolSize)

	// Get strength
	strength := entropy.StrengthEvaluation(total)

	fmt.Println("Password:", password)
	fmt.Printf("Entropy: %.2f bits\n", total)
	fmt.Println("Strength:", strength)
}