package exercices

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	n := 44
	expected := 701408733
	result := Fibonacci(n)
	if result != expected {
		t.Fatalf("want %v, got %v", expected, result)
	}

}
