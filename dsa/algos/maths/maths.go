package maths

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
  a=bq+r
*/

// gcd(a,b)== gcd(b,r); Need to divide until one of the number is zero.
func Gcd(a uint, b uint) uint {
	if a == 0 {
		return b
	}
	return Gcd(b%a, a)
}

// Find gcd of array numbers pair wise.
func GcdArr(a []uint) uint {
	gcd := a[0]
	for i := 1; i < len(a); i++ {
		gcd = Gcd(gcd, a[i])
	}
	return gcd
}

// a*b=lcm(a,b) *gcd(a,b)
func lcm(a uint, b uint) uint {
	return (a * b) / Gcd(a, b)
}

func lcmArray(a []int) int {
	cumLcm := a[0]
	for i := 1; i < len(a); i++ {
		cumLcm = int(lcm(uint(cumLcm), uint(a[i])))
	}
	return int(cumLcm)
}

func isPrime(number int) bool {
	sqrt := math.Sqrt(float64(number))
	for i := 2; i <= int(sqrt); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func FactorsArr(x int) []int {
	if x > 0 {
		factors := make([]int, 0, x/2)
		factors = append(factors, 1)
		for i := 2; i <= x/2; i++ {
			if x%i == 0 {
				factors = append(factors, i)
			}
		}
		factors = append(factors, x)
		return factors
	}
	return nil
}

func Factors(x int) string {
	if x > 0 {
		var sb strings.Builder
		sb.WriteString("1, ")
		for i := 2; i <= x/2; i++ {
			if x%i == 0 {
				sb.WriteString(strconv.Itoa(i) + ", ")
			}
		}
		sb.WriteString(strconv.Itoa(x))
		return sb.String()
	}
	return ""
}

func PrimeFactors(n int) string {
	// Print the number of 2s that divide n
	fmt.Printf("Prime factor of= %v => ", n)
	for n%2 == 0 {
		fmt.Print("2 ")
		n /= 2
	}

	// n must be odd at this point.  So we can
	// skip one element (Note i = i +2)
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		// While i divides n, print i and divide n
		for n%i == 0 {
			fmt.Print(i)
			n /= i
		}
	}

	// This condition is to handle the case when
	// n is a prime number greater than 2
	if n > 2 {
		fmt.Print(n)
	}
	return ""

}

// Exponentiation
// O(n)
func PowerR(x int, exponent int) int {
	if exponent == 0 {
		return 1
	} else {
		return x * PowerR(x, exponent-1)
	}
}

func PowerRBinary(x int, exponent int) int {
	if exponent == 0 {
		return 1
	} else if exponent%2 == 0 {
		return PowerR(x*x, exponent/2)
	} else {
		return x * PowerR(x*x, (exponent-1)/2)
	}
}
