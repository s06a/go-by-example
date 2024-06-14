package main

import "fmt"

// fact function returns factorial of a number
func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// recursiveSum returns sum of numbers from 1 to n
func recursiveSum(n int) int {
	if n == 0 {
		return 0
	}
	return n + recursiveSum(n-1)
}

func main() {
	// call factorial function
	fmt.Println(fact(7))

	// call fibonacci function
	fmt.Println(fibonacci(7))

	// call recursiveSum function
	fmt.Println(recursiveSum(5))

	// closures can also be recursive but this requires
	// to be declared with a typed var explicitly before
	// it's defined
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	// calling fib func
	fmt.Println(fib(7))
}
