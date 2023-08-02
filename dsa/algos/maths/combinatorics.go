package maths

import (
	"fmt"
)

/**
 * Counting.
Basic counting rules.

Product rule - Suppose that a procedure can be broken down into a sequence of two tasks. If there are n1 ways to do the first task and for each of these ways of doing the first task, there are n2 ways to do the second task, then there are n1n2 ways to do the procedure.<br> </br>
Sum rule - If a task can be done either in one of n1 ways or in one of n2 ways, where none of the set of n1 ways is the same as any of the set of n2 ways, then there are n1 +n2 ways to do the task. <br> </br>
Subtraction rule : If a task can be done in either n1 ways or n2 ways, then the number of ways to do the task is n1 + n2 minus the number of ways to do the task that are common to the two different ways. <br> </br>
Division Rule - There are n/d ways to do a task if it can be done using a procedure that can be carried out in n ways, and for every way w, exactly d of the n ways correspond to way w. </br>

Permutations of {1, . . . , n} is a necessary prerequisite to generating them.
   		Permutations of n elements fact(n)= n!=n * (n-1) * (n-2) … 1


Permutation of r item of n :
A permutation of a set of distinct objects is an ordered arrangement of these objects. We also are interested in ordered arrangements of some of the elements of a set. An ordered arrangement of r elements of a set is called an r-permutation.

P (n, r ) = n (n − 1) (n − 2) · · ·  (n − r + 1)   => P(n, r)= n! / (n-r)! </br>

Combination :
The number of r-combinations of a set with n elements, where n is a nonnegative integer and r is an integer with 0 ≤ r ≤ n, equals </br>

  C(n, r) = n!/(r! * (n-r)!)

  </br>

  Subsets of an set of size n = pow(2,n). If n=3 then subsets=8    </br>



  Multiset Counting : 5-Red, 3-Yellow , 2-While => Permutations = (5+3+2)! / 5 !  3! 2!

Strings over finite alphabet : The number of k-digit strings over an n-element alphabet is pow(2,k).
         n=2 , 0,1     k=2 => pow(2,2)=4      k=3  => pow(2,3)= 8.

 The number of all subsets of an n-element set is 2n.

*/

/**
 *   factorial(0) =1
 *
 *   Arranging n distinct objects. n * n-1, n-2......1
 *   a,b,c => abc, acb, bac, bca, cab, cba  == 6  => 3!=6
 *
 * @param x
 * @return
 */
func Factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}

/*
*
  - The number of r-combinations of a set with n elements, where n is a nonnegative integer and r is an integer with 0 ≤ r ≤ n, equals
  - C(n, r) = n!/(r! * (n-r)!)
  - C(n, r) = C(n-1, r-1) + C(n-1, r)
    C(n, 0) = C(n, n) = 1
  - @param n
  - @param r
  - @return
*/
func Ncr(n int, r int) int {
	if r == 0 || n == r {
		return 1
	}
	combinations := Factorial(n) / (Factorial(n-r) * Factorial(r))
	return combinations
}

/**
 *   A permutation of a set of distinct objects is an ordered arrangement of these objects.
 *   We also are interested in ordered arrangements of some of the elements of a set.
 *   An ordered arrangement of r elements of a set is called an r-permutation.
 *
 * P (n, r ) = n (n − 1) (n − 2) · · ·  (n − r + 1)   => P(nor)= n! / (n-r)!
 * @return
 */
func Npr(n int, r int) int {
	permuations := Factorial(n) / Factorial(n-r)
	return permuations

}

/**
 * power((x+y),n)
 */
func binomialTheorem(binomialSize int) {
	fmt.Printf("power((x+y),%v) => ", binomialSize)
	for i := binomialSize; i >= 0; i-- {
		coeff := Ncr(binomialSize, i)
		fmt.Printf("%v * power(x,%v) * power(x,%v) + ", coeff, i, binomialSize-i)
	}
}
