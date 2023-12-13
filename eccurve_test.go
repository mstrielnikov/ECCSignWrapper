package main

import (
	"crypto/elliptic"
	"fmt"
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
	G := Curve.BasePointGGet()

	sum := Curve.AddECPoints(G, G)

	double := Curve.DoubleECPoints(G)

	fmt.Println(sum)
	fmt.Println(double)

	if !Curve.IsOnCurveCheck(sum) {
		t.Fatalf("(A + A) point is not on the curve")
	}

	if !Curve.IsOnCurveCheck(double) {
		t.Fatalf("(2A) point is not on the curve")
	}

	if !sum.IsEqual(&double) {
		t.Fatalf("ECC doubling addition test failed")
	}
}
