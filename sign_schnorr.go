package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// SchnorrSignature represents a Schnorr signature.
type SchnorrSignature struct {
	R, S *big.Int
}

// GenerateKeyPair generates a key pair for Schnorr signatures.
func (ec *ECCurve) GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	priv, err := ecdsa.GenerateKey(ec.curve, rand.Reader)

	if !ec.IsOnCurveCheck(ECPoint{priv.X, priv.Y}) {
		return nil, nil, fmt.Errorf("private key point in not on the curve")
	}

	if err != nil {
		return nil, nil, err
	}
	return priv, &priv.PublicKey, nil
}

// Sign generates a Schnorr signature for a given message and private key.
func (ec *ECCurve) Sign(message *big.Int, priv *ecdsa.PrivateKey) (*SchnorrSignature, error) {
	// Generate a random nonce
	k, err := rand.Int(rand.Reader, ec.curve.Params().N)
	if err != nil {
		return nil, err
	}

	// Calculate R = k * G
	Rx, _ := ec.curve.ScalarBaseMult(k.Bytes())

	// Calculate r = H(R || pubKey || message)
	r := hashPoints(Rx, priv.PublicKey.X, priv.PublicKey.Y, message)

	// Calculate s = k + e * privateKey
	s := new(big.Int).Mul(r, priv.D)
	s.Add(s, k)
	s.Mod(s, ec.curve.Params().N)

	// The signature is (R, s)
	return &SchnorrSignature{R: Rx, S: s}, nil
}

// Verify verifies a Schnorr signature for a given message and public key.
func (ec *ECCurve) Verify(signature *SchnorrSignature, message *big.Int, pub *ecdsa.PublicKey) bool {
	// Check for nil public key
	if pub == nil {
		return false
	}

	// Calculate e = H(R || pubKey || message)
	e := hashPoints(signature.R, pub.X, pub.Y, message)

	// Calculate R' = s * G - e * pubKey
	var Rx, Ry *big.Int
	Rx, Ry = ec.curve.ScalarBaseMult(signature.S.Bytes())

	// Check for nil curve points
	if signature.R == nil {
		return false
	}

	tempX, tempY := ec.curve.ScalarMult(pub.X, pub.Y, e.Bytes())
	Rx, Ry = ec.curve.Add(Rx, Ry, tempX, new(big.Int).Neg(tempY))

	// Check if R == R'
	return signature.R.Cmp(Rx) == 0 && Ry != nil
}

// hashPoints concatenates the coordinates of given points and the message, and then hashes the result using SHA-256.
func hashPoints(points ...*big.Int) *big.Int {
	hash := sha256.New()
	for _, point := range points {
		if point != nil {
			hash.Write(point.Bytes())
		}
	}
	return new(big.Int).SetBytes(hash.Sum(nil))
}
