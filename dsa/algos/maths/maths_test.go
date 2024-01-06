package maths

import (
	"testing"
)

func TestGcd(t *testing.T) {
	var a uint = 10
	var b uint = 15
	input_arr := [3]uint{10, 15, 25}
	var expected uint = 5
	gcd := Gcd(a, b)
	if gcd != expected {
		t.Fatalf("Error in Gcd() Expected= %v => actual = %v", expected, gcd)
	}
	gcd_arr := GcdArr(input_arr[:])
	if gcd_arr != expected {
		t.Fatalf("Error in Gcd() Expected= %v => actual = %v", expected, gcd_arr)
	} else {
		t.Logf("Gcd() Expected= %v => actual = %v", expected, gcd)
	}
}

func TestSummations(t *testing.T) {
	//fmt.Printf("\n Sum1toN(%v) = %v ", 10, Sum1toN(10))
	Sum1toN(10)
	SumOfSquares(4)
	SumOfGeometric(1, 2, 3)

}
