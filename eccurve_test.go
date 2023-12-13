package main

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"testing"
)

func TestECCAlgebraEquation(t *testing.T) {
	curveParam := elliptic.P256()
	Curve := NewECCurve(curveParam)

	G := Curve.BasePointGGet()
	k := SetRandom(256)
	d := SetRandom(256)

	H1 := Curve.ScalarMult(d, G)
	H2 := Curve.ScalarMult(k, H1)

	H3 := Curve.ScalarMult(k, G)
	H4 := Curve.ScalarMult(d, H3)

	result := H2.IsEqual(&H4)

	fmt.Println(H2)
	fmt.Println(H4)

	if !Curve.IsOnCurveCheck(H2) {
		t.Fatalf("H2 is not on the curve")
	}

	if !Curve.IsOnCurveCheck(H4) {
		t.Fatalf("H4 is not on the curve")
	}

	if !result {
		t.Fatalf("ECC Algebra Equation test failed")
	}
}

func TestAdditionDoubling(t *testing.T) {
	curveParam := elliptic.P256()
	Curve := NewECCurve(curveParam)

	pointA := ECPoint{X: big.NewInt(1), Y: big.NewInt(2)}

	sum := Curve.AddECPoints(pointA, pointA)

	double := Curve.DoubleECPoints(pointA)

	if !sum.IsEqual(&double) {
		t.Fatalf("ECC doubling addition test failed")
	}
}
