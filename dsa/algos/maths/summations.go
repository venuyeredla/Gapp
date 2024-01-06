package maths

import (
	"fmt"
	"math"
)

/**
 *  sum{1.. n} = (n*(n+1))/2
 */
func Sum1toN(upto int) {
	sum := 0
	for i := 1; i <= upto; i++ {
		sum = sum + i
	}
	fmt.Printf("Iterative sum : %v ", sum)
	sum = (upto * (upto + 1)) / 2
	fmt.Printf("\n%v =>  (upto * (upto+1))/2 = %v", upto, sum)
}

// s= (n*(n+1)*(2n+1))/6
func SumOfSquares(upto int) {
	sum := (upto * (upto + 1) * (2*upto + 1)) / 6
	fmt.Printf("\n %v => (upto * (upto+1) * (2*upto+1))/2 = %v ", upto, sum)
}

/**
 *  if r!=1   sum= (a * pow(r,n+1) - 1)/r-1
 *  if a=1 : sum = (pow(r,n+1) - 1)/r-1
 *
 * @param upto
 */
func SumOfGeometric(a int, r int, n int) {
	power := int(math.Pow(float64(r), float64(n+1)))
	dividend := (a * power) - a
	sum := dividend / (r - 1)
	fmt.Printf("\n %v => (a * (int)Math.pow(r,n+1) - 1)/r-1 = %v ", n, sum)
}
