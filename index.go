package main

import (
	"fmt"

	"github.com/Samy0412/gopractice/exercices"
)

func main() {
	fmt.Println("Exercice 1 :")
	fmt.Println(exercices.Split("I have a cat and a dog"))
	fmt.Println("Exercice 2 :")
	fmt.Println(exercices.MoreThanFiveLetters([]string{"dog", "elephant", "cat", "hippopotamus", "giraffe"}))
}
