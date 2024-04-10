package main

import "testing"

// Dummy function: Sum two integers
func SumInt(number1 int, number2 int) int {
	return number1 + number2
}

// Test SumInt function
func Test_SumInt_RangeOfValues(t *testing.T) {
	// Test table struct
	type testType struct {
		name    string
		number1 int
		number2 int
		want    int
	}

	// Test table cases
	testCases := []testType{
		{"odd plus odd", 1, 1, 2},
		{"odd plus even", 1, 2, 3},
		{"even plus even", 2, 2, 4},
		{"one zero", 0, 2, 2},
		{"both zero", 0, 0, 0},
		{"one negative", -3, 5, 2},
		{"both negative", -3, -2, -5},
		{"big numbers", 10000, 10000, 20000},
	}

	// Test each scenario as a sub-test
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SumInt(tc.number1, tc.number2)
			if got != tc.want {
				t.Errorf("For (%d + %d) --> want: %d, got: %d\n", tc.number1, tc.number2, tc.want, got)
			}
		})
	}
}
