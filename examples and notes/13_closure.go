package main

// closures are functions that can access or manipulate
// values from the outer function. also, here we use
// anonymous function, which is useful when you want to
// define a function without a name.

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// factorial function returns a function which gets
// integer and returns integer. first we define result
// as a value in the outer function, then we return the
// inner function which uses result as a value.
func factorial() func(int) int {
	var result int
	return func(num int) int {
		if num <= 1 {
			return 1
		}
		result = 1
		for i := 2; i <= num; i++ {
			result *= i
		}
		return result
	}
}

// this function gets factor and returns a function
// which gets integer (num) and returns integer.
func multiplyBy(factor int) func(int) int {
	return func(num int) int {
		return factor * num
	}
}

func main() {
	// next integer
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newNextInt := intSeq()
	fmt.Println(newNextInt())

	// factorial function
	returnFactorial := factorial()

	fmt.Println(returnFactorial(5))
	fmt.Println(returnFactorial(10))

	// multiplyBy function
	multiplyByTwo := multiplyBy()
	fmt.Println(multiplyByTwo(5))
}
