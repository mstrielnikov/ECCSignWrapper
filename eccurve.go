package main

import (
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"math/big"
)

// ECPoint represents a point on an elliptic curve.
type ECPoint struct {
	X, Y *big.Int
}

// ECPointGen creates a new ECPoint from given coordinates.
func ECPointGen(x, y *big.Int) ECPoint {
	return ECPoint{X: x, Y: y}
}

// IsEqual checks if two elliptic curve points are equal.
func (ecp *ECPoint) IsEqual(ecPoint *ECPoint) bool {
	return ecp.X.Cmp(ecPoint.X) == 0 && ecp.Y.Cmp(ecPoint.Y) == 0
}

// ECCurve implements methods for Edwards curve operations.
type ECCurve struct {
	curve elliptic.Curve
}

// NewECCurve initializes a new ECCurve.
func NewECCurve(curve elliptic.Curve) *ECCurve {
	return &ECCurve{curve: curve}
}

// BasePointGGet returns the base point (generator) of the curve.
func (ec *ECCurve) BasePointGGet() ECPoint {
	x, y := ec.curve.Params().Gx, ec.curve.Params().Gy
	return ECPoint{X: x, Y: y}
}

// ScalarMult multiplies a point by a scalar and returns the result.
func (ec *ECCurve) ScalarMult(scalar *big.Int, point ECPoint) ECPoint {
	x, y := ec.curve.ScalarMult(point.X, point.Y, scalar.Bytes())
	return ECPoint{X: x, Y: y}
}

// IsOnCurveCheck checks if a given point is on the curve.
func (ec *ECCurve) IsOnCurveCheck(point ECPoint) bool {
	return ec.curve.IsOnCurve(point.X, point.Y)
}

// AddECPoints adds two elliptic curve points and returns the result.
func (ec *ECCurve) AddECPoints(pointA, pointB ECPoint) ECPoint {
	x, y := ec.curve.Add(pointA.X, pointA.Y, pointB.X, pointB.Y)
	return ECPoint{X: x, Y: y}
}

// DoubleECPoints doubles the given elliptic curve point and returns the result.
func (ec *ECCurve) DoubleECPoints(point ECPoint) ECPoint {
	x, y := ec.curve.Double(point.X, point.Y)
	return ECPoint{X: x, Y: y}
}

// ECPointToString serializes an ECPoint to a JSON-encoded string.
func ECPointToString(point ECPoint) string {
	data, _ := json.Marshal(point)
	return string(data)
}

// StringToECPoint deserializes an ECPoint from a JSON-encoded string.
func StringToECPoint(s string) ECPoint {
	var point ECPoint
	_ = json.Unmarshal([]byte(s), &point)
	return point
}

// PrintECPoint prints the coordinates of an ECPoint.
func PrintECPoint(point ECPoint) {
	fmt.Printf("X: %s\nY: %s\n", point.X.String(), point.Y.String())
}
