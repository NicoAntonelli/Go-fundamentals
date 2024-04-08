package main

import "fmt"

// Work with arrays, slices and maps
func arraySliceMap() {
	// Arrays
	fmt.Println("Arrays")
	var names [3]string = [3]string{"John", "Jane", "Jack"}
	namesAgain := [...]string{"John", "Jane", "Jack"}
	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(namesAgain[1:3])
	fmt.Println()
	
	// Slices
	fmt.Println("Slices")
	var animals = []string{"cat", "dog", "fish"}
	fmt.Println(animals)
	fmt.Printf("The length is %v with capability of %v\n", len(animals), cap(animals))
	animals = append(animals, "bird")
	fmt.Println(animals)
	fmt.Printf("The length is %v with capability of %v\n", len(animals), cap(animals))
	var animals2 = []string{"lion", "tiger", "bear"}
	animals = append(animals, animals2...)
	fmt.Println(animals)
	fmt.Printf("The length is %v with capability of %v\n", len(animals), cap(animals))
	var numbers []int32 = make([]int32, 3, 5)
	fmt.Println(numbers)
	fmt.Printf("The length is %v with capability of %v\n", len(numbers), cap(numbers))
	fmt.Println()

	// Maps
	fmt.Println("Maps")
	var students1 map[string]uint8 = make(map[string]uint8)
	students1["John"] = 9
	students1["Mark"] = 6
	delete(students1, "Mark")
	fmt.Println(students1["John"])
	fmt.Println(students1["Mark"])
	
	students2 := map[string]uint8 {"Pepe": 8, "Jack": 7}
	var grade, ok = students2["Pepe"]
	if (ok) {
		fmt.Println(grade)
	} else {
		fmt.Println("Student not found")
	}
	fmt.Println()
}
