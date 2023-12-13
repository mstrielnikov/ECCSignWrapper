package main

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"testing"
)

func TestSchnorSign(t *testing.T) {
	curve := elliptic.P256()
	ellipticCurve := NewECCurve(curve)

	// Generate key pair
	privKey, pubKey, err := ellipticCurve.GenerateKeyPair()
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Convert []byte message to *big.Int
	message := new(big.Int).SetBytes([]byte("Hello, Schnorr!"))
	fmt.Println("Message as big int:", message)

	// Sign a message
	signature, err := ellipticCurve.Sign(message, privKey)
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println("Signature:", &signature)

	// Verify the signature
	verified := ellipticCurve.Verify(signature, message, pubKey)
	if verified {
		fmt.Println("Signature is valid.")
	} else {
		t.Fatalf("Signature is invalid.")
	}
}
