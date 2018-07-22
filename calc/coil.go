package calc

import "math"

// MicrostripInductance returns the inductance of a microstrip in uH.
// Capacitance is likely to be quite high.
//
// Formula source:
//   Lumped Elements for RF and Microwave Circuits - Inder Bahl 2003
//   http://www.lmn.pub.ro/~daniel/ElectromagneticModelingDoctoral/Books/EM%20Field/Bahl2003-Lumped%2520Elements%2520for%2520RF%2520and%2520Microwave%2520Circuits.pdf
func MicrostripInductance(mmLength, mmWidth float64) float64 {
	c := math.Log(4*mmLength/mmWidth) + (0.0224 * mmWidth / (2 * mmLength)) + 0.5
	return 0.0002 * mmLength * c
}

// Source: http://coil32.net/pcb-coil.html
func fillFactor(umOuterDia, umInnerDia float64) float64 {
	return (umOuterDia - umInnerDia) / (umOuterDia + umInnerDia)
}

// Source: http://coil32.net/pcb-coil.html
func avgDiameter(umOuterDia, umInnerDia float64) float64 {
	return (umOuterDia + umInnerDia) / 2
}

// SquareCoilinductance calculates the inductance of a square coil.
// Returns the inductance in nano-henrys.
//
// L = K1 * U0 * ((N^2 * Davg)/(1 + K2*fillFactor))
// K1,K2 are contants based on the shape of the inductor (square)
// U0 is the magnetic constant (4pi * 10^-7)
//
// Formula source:
//    http://coil32.net/pcb-coil.html
func SquareCoilinductance(numTurns float64, umOuterDia, umInnerDia float64) float64 {
	t := numTurns * numTurns * avgDiameter(umOuterDia, umInnerDia)
	b := (2.75 * fillFactor(umOuterDia, umInnerDia)) + 1
	return 2.34 * 4 * math.Pi * 10e-07 * (t / b)
}

// RoundCoilInductance calculates the inductance of a circular coil.
// Returns the inductance in nano-henrys.
//
// Formula source:
//    http://coil32.net/pcb-coil.html
func RoundCoilInductance(numTurns float64, umOuterDia, umInnerDia float64) float64 {
	o1 := (4 * math.Pi * 10e-07 * numTurns * numTurns * avgDiameter(umOuterDia, umInnerDia)) / 2
	ff := fillFactor(umOuterDia, umInnerDia)
	o2 := math.Log(2.46/ff) + (0.20 * ff * ff)
	return o1 * o2
}
