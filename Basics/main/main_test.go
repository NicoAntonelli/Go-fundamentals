package main

import "testing"

// Dummy function
func Hello() string {
    return "Hello world"
}

// Test Hello function
func TestHello(t *testing.T) {
    want := "Hello world"
	got := Hello();
    if (got != want) {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
