package main

import (
	"crypto/rand"
	"math/big"
)

func SetRandom(bits int) *big.Int {
	// Generate a random number with the specified number of bits
	randomBits, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		// Handle the error, e.g., log it or return a default value
		panic(err)
	}

	// Create a new big.Int and set its value to the generated random number
	randomInt := new(big.Int).Set(randomBits)

	return randomInt
}
