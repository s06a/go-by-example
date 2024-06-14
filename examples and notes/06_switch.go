package main

import (
	"fmt"
	"time"
)

func main() {

	switch time.Now().Weekday() {
	case time.Thursday, time.Friday:
		fmt.Println("get back to sleep")
	default:
		fmt.Println("go to work!")
	}

	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	checkType := func(i interface()) {
		// interface is a special type that can hold
		// vallues of any type (int, bool, and ...)
		switch t := i.(type) {
		case bool:
			fmt.Println("it's boolean")
		case int:
			fmt.Println("it's integer")
		default:
			fmt.Printf("don't know %T", t)
		}
	}

	checkType(bool)
	checkType(1)
	checkType("hey")

}
