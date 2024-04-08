package main

import "fmt"

// Work with loops using FOR
func forLoop() {
	// Simple for loop
	fmt.Println("For loop")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println()

	// For loop as a while loop
	fmt.Println("For loop as a while loop")
	otherSum := 0
	index := 0
	for otherSum < 1000 {
		otherSum += index
		index++
	}
	fmt.Println(sum)
	fmt.Println()
	
	// For loop handled only with break
	fmt.Println("For loop handled only with break")
	index = 0
	for {
		if (index >= 4) {
			break
		}
		index++
		fmt.Println("Index:", index)
	}
	fmt.Println()

	// Iteration over array/slice
	var animals = []string{"lion", "tiger", "bear"}
	fmt.Println("Iteration over array/slice")
	for i, v := range animals {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
	fmt.Println()
	
	// Iteration over a map
	students := map[string]uint8 {"Pepe": 8, "Jack": 7}
	fmt.Println("Iteration over a map")
	for name, grade := range students {
		fmt.Printf("Name: %v, Grade: %v\n", name, grade)
	}
	fmt.Println()
}
