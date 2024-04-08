package main

import "fmt"

// Work with structs
func structs() {
	fmt.Println("Structs")
	var student1 Student = Student{name: "Nick", age: 26}
	student1.subjects = []string{"Math", "English", "Dev"}
	student1.University = University{"UTN"}
	student1.univName = "UTN"

	var student2 Student = Student{"Herno", 25, []string{"Math", "English", "Dev"}, nil, University{"UTN"}}
	student2.grades = []Grade{{"Math", 9}, {"English", 8}}
	student2.grades = append(student2.grades, Grade{"Dev", 8})
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println("Average grade of student 2:", student2.getAverageGrade())
	fmt.Println()

	// Structs with interface
	fmt.Println("Struct with inteface")
	var engine GasEngine = GasEngine{model: "V8", liters: 50, kmPerLiter: 10}
	fmt.Printf("Engine %v has %v kilometers left\n", engine.model, engine.kilometersLeft())
	canMakeIt(engine, 100)
	var newEngine ElectricEngine = ElectricEngine{model: "TSLA", kwh: 20, kmPerKwh: 20}
	fmt.Printf("New Engine %v has %v kilometers left\n", newEngine.model, newEngine.kilometersLeft())
	canMakeIt(newEngine, 100)
	var otherEngine Engine = GasEngine{liters: 10, kmPerLiter: 8}
	fmt.Printf("The other engine has %v kilometers left\n", otherEngine.kilometersLeft())
	canMakeIt(otherEngine, 100)
	fmt.Println()
}
