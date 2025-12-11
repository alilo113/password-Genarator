package entropy

import (
	"math"
	"pwman/internal/generator"
)

// charSizeSet returns the pool size based on character types used in the password
func CharSizeSet(password string) int {
	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSymbol := false

	for _, c := range password {
		switch {
		case c >= '0' && c <= '9':
			hasDigit = true
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		default:
			hasSymbol = true
		}
	}

	poolSize := 0
	if hasLower {
		poolSize += 26
	}
	if hasUpper {
		poolSize += 26
	}
	if hasDigit {
		poolSize += 10
	}
	if hasSymbol {
		poolSize += 32 // match your generator's symbols
	}

	return poolSize
}

// EntropyPerChar calculates bits of entropy per character
func EntropyPerChar(poolSize int) float64 {
	if poolSize <= 0 {
		return 0
	}
	return math.Log2(float64(poolSize))
}

// TotalEntropy calculates total bits of entropy for the password
func TotalEntropy(password string, poolSize int) float64 {
	bitsPerChar := EntropyPerChar(poolSize)
	return bitsPerChar * float64(len(password))
}

// StrengthEvaluation returns a human-readable label based on total entropy
func StrengthEvaluation(totalEntropy float64) string {
	switch {
	case totalEntropy < 40:
		return "very weak"
	case totalEntropy < 60:
		return "weak"
	case totalEntropy < 80:
		return "medium"
	case totalEntropy < 100:
		return "strong"
	default:
		return "extremely strong"
	}
}

// Example usage: calculate strength of generated password
func Example() {
	// Get a password from generator
	password := generator.GeneratePassword(20)

	// Calculate pool size
	poolSize := CharSizeSet(password)

	// Calculate total entropy
	total := TotalEntropy(password, poolSize)

	// Get strength label
	strength := StrengthEvaluation(total)

	// Output
	println("Password:", password)
	println("Entropy (bits):", total)
	println("Strength:", strength)
}