package main

import "time"

// Measure the execution time of a loop
func testLoopTime(slice []int, num int) time.Duration {
	var timeBase = time.Now()
	for len(slice) < num {
		slice = append(slice, 1)
	}
	return time.Since(timeBase)
}
