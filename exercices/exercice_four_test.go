package exercices

import (
	"testing"
)

func TestFibonacciMemoized(t *testing.T) {
	n := 44
	expected := 701408733
	result := FibonacciMemoized(n)
	if result != expected {
		t.Fatalf("want %v, got %v", expected, result)
	}

}
