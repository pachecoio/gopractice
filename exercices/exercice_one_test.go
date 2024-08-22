package exercices

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	x := "I have a cat and a dog"
	expected := []string{"I", "have", "a", "cat", "and", "a", "dog"}
	result := Split(x)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("want %v, got %v", expected, result)
	}
}

func TestSplitEmpty(t *testing.T) {
	x := ""
	expected := []string{""}
	result := Split(x)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("want %v, got %v", expected, result)
	}
}
