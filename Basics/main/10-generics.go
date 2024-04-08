package main

import "fmt"

type Car struct {
	carMake string
	carModel string
}

type Motorcycle struct {
	motorMake string
	motorModel string
}

type Pilot[T Car | Motorcycle] struct {
	name string
	age int
	vehicle T
}

// Work with generics
func generics() {
	fmt.Println("Generics")
	var intSlice = []int{1, 2, 3, 4, 5}
	var float32Slice = []float32{1.2, 2.3, 3.4, 4.5, 5.6}
	var float64Slice = []float64{1.2, 2.3, 3.4, 4.5, 5.6}
	fmt.Println("Int sum:", sumGeneric[int](intSlice))
	fmt.Println("Float32 sum:", sumGeneric[float32](float32Slice))
	fmt.Println("Float64 sum:", sumGeneric[float64](float64Slice))
	fmt.Println("Float64 isEmpty:", isEmpty(float64Slice))
	fmt.Println()

	fmt.Println("Struct with generics")
	var pilot1 = Pilot[Car]{
		name: "John",
		age: 30,
		vehicle: Car{
			carMake: "Ford",
			carModel: "Mustang",
		},
	}

	var pilot2 = Pilot[Motorcycle]{
		name: "Jack",
		age: 25,
		vehicle: Motorcycle{
			motorMake: "Honda",
			motorModel: "CBR",
		},
	}

	fmt.Printf("Pilot 1: %+v\n", pilot1)
	fmt.Printf("Pilot 2: %+v\n", pilot2)
	fmt.Println()
}

// Generic sum
func sumGeneric[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Generic isEmpty check
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
