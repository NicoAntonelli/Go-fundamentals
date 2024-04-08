package main

import (
	"fmt"
	"unicode/utf8"
)

// Practice code
func main() {
	// Hello World
	text := "Hello world"
	fmt.Println(text)
	var counter int = stringLen(text)
	fmt.Println("The previous sentence has", counter, "characters")
	fmt.Printf("The counter %v is of type %T\n", counter, counter)
	fmt.Println()

	// Function with validation
	funcValidation()
	
	// Array, slice & map
	arraySliceMap()
	
	// For loop
	forLoop()
	
	// Test time for a loop
	fmt.Println("Testing time")
	var num int = 1000000
	var testSlice1 = []int{}
	var testSlice2 = make([]int, 0, num)
	fmt.Printf("Total time without preallocation: %v\n", testLoopTime(testSlice1, num))
	fmt.Printf("Total time with preallocation: %v\n", testLoopTime(testSlice2, num))
	fmt.Println()
	
	// String vs runes
	stringRunes()

	// Structs
	structs()

	// Pointers
	pointers()

	// Go Routines with Wait Groups and Mutex
	goRoutinesWaitGroups()

	// Go Routines with Channels
	goRoutinesChannels()

	// Generics
	generics()

	// Files and generics
	files()

	// Handling errors
	handlingErrors()
}

// Length of a string including non-ascii chars
func stringLen(text string) int {
	return utf8.RuneCountInString(text)
}
