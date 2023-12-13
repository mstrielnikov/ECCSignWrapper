package main

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"testing"
)

func TestECCAlgebraEquation(t *testing.T) {
	var G ECPoint
	var curveParam elliptic.Curve
	var k, d *big.Int

	curveParam = elliptic.P256()
	Curve := NewECCurve(curveParam)

	G = Curve.BasePointGGet()
	k = SetRandom(256)
	d = SetRandom(256)

	H1 := Curve.ScalarMult(d, G)
	H2 := Curve.ScalarMult(k, H1)

	H3 := Curve.ScalarMult(k, G)
	H4 := Curve.ScalarMult(d, H3)

	result := H2.IsEqual(&H4)

	fmt.Println(H2)
	fmt.Println(H4)

	if !result {
		t.Fatalf("ECC Algebra Equation test failed")
	}
}
