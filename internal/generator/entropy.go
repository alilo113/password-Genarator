package main

import (
    "fmt"
    "pwman/internal/generator"
	"math"
)

func charSizeSet() {
    password := generator.GeneratePassword()
    chars := []rune(password)

    hasLower := false
    hasUpper := false
    hasDigit := false
    hasSymbol := false

    for _, c := range chars {
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
        poolSize += 32 // or whatever your symbol count actually is
    }
}

func entropyPerChar(poolSize int) float64 {
	return math.Log2(float64(poolSize))
}