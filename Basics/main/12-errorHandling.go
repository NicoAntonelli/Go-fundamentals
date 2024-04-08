package main

import (
	"errors"
	"fmt"
)

// Handling errors with panic, recover and defer
func handlingErrors() {
	fmt.Println("Error handling")
	err := perform()
	fmt.Println(err)
	fmt.Println()
}

// Execute some logic and handle a possible error
func perform() (err error) {
	// Set anonymous function and execute it in the end
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	// Call a function with an error
	goesWrong()
	return
}

// Function that throws an error
func goesWrong() {
	panic(errors.New("show fail message"))
}
