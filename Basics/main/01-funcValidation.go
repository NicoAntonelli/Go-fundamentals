package main

import (
	"errors"
	"fmt"
)

// Errors validation
func funcValidation() {
	fmt.Println("Function with validation")
	var numerator, denominator int = 5, 2
	fmt.Printf("Calculate %v/%v: ", numerator, denominator)
	var result, reminder, exception = intDivision(numerator, denominator)

	// Guard clause
	if (exception != nil) {
		fmt.Println(exception.Error())
		return
	}

	if (reminder != 0) {
		fmt.Printf("The result is %v and the reminder is %v\n", result, reminder)
	} else {
		fmt.Printf("The result is %v\n", result)
	}
	
	fmt.Println()
}

// Integer division
func intDivision(numerator, denominator int) (int, int, error) {
	var Exception error // nil
	if (denominator == 0) {
		Exception = errors.New("denominator cannot be zero")
		return 0, 0, Exception
	}

	var result = numerator / denominator
	var reminder = numerator % denominator
	return result, reminder, Exception
}
