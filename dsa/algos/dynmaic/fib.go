package dynmaic

import "fmt"

// fib(0)=0,fib(1)=1
// f(n)=f(n-1)+fib(n-2) return nth term in sequence.
func fib(n int) int {
	if n == 0 || n == 1 {
		fmt.Printf("fib(%v)", n)
		return n
	} else {
		fmt.Printf("fib(%v) ,fib(%v)\n", n-1, n-2)
		return fib(n-1) + fib(n-2)
	}
}
