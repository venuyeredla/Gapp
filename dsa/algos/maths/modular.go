package maths

import "fmt"

/**
Let m be a positive integer and let a and b be integers. <br>
* a is congruent to b under module m .     m | a-b  <=> a%m=b%m=r. Same reminder.<br>
* a-b=mk => a=mk+b;<br>

 If a ≡ b (mod m) and c ≡ d (mod m),
 then 1. a + c ≡ b + d (mod m)
      2. ac ≡ bd (mod m).<br><br>
      3. (a + b)%m = ((a%m) + (b%m))%m
      4. (a*b)% m = ((a%m)*(b%m))%m.
      5. (a-b)%m = ((a%m)-(b%m)+m)%m
	  6. (a/b)%m =((a%m)*((1/b)%m))%m

	Extended Eucledan algorithm : gcd(a,b)=ax+by - linear combination.
	Coefficients x and y are used to find modular inverse.
*/

func IsCongruent(a int, b int, modulo int) bool {
	fmt.Printf(" a mod m = b mod m => %v = %v ", a%modulo, b%modulo)
	return (a-b)%modulo == 0
}

// a  +(mod m) b = a+b (mod m)
func ModAdd(input []int, modulo int) int {
	modSum := 0
	for _, val := range input {
		modSum = (modSum + val) % modulo
	}
	return modSum
}

// a  *(mod m) b = a*b (mod m)
func ModMulti(input []int, modulo int) int {
	modProduct := 1
	for _, val := range input {
		modProduct = (modProduct * val) % modulo
	}
	return modProduct
}

/**
 *   power(5,25) (mod 4)
 */

func ModExponent(base int, exponent int, modulo int) int {
	result := 1
	powerToCarry := base % modulo
	for exponent != 0 {
		number := exponent & 1
		fmt.Printf("Binary digit : %v", number)
		if number == 1 {
			result = (result * powerToCarry) % modulo
			fmt.Printf("result : %v \n", result)
		}
		powerToCarry = (powerToCarry * powerToCarry) % modulo
		fmt.Printf("Power to carry : %v", powerToCarry)
		exponent = exponent >> 1
	}
	return result
}

/**
 * Modular exponentiation .
 */
func ModuloPower(x int, y int, p int) int {
	res := 1 // Initialize result
	for y > 0 {
		// If y is odd, multiply x with result
		if (y & 1) != 0 {
			res = res * x
		}
		// y must be even now
		y = y >> 1 // y = y/2
		x = x * x  // Change x to x^2
	}
	return res % p
}
