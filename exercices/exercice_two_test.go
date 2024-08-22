package exercices

import (
	"reflect"
	"testing"
)

func TestMoreThanFiveLetters(t *testing.T) {
	x := []string{"dog", "elephant", "cat", "hippopotamus", "giraffe"}
	expected := []string{"elephant", "hippopotamus", "giraffe"}
	result := MoreThanFiveLetters(x)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("want %v, got %v", expected, result)
	}
}
