package main

import (
	"fmt"

	"github.com/Samy0412/gopractice/exercices"
)

func main() {
	// fmt.Println("Exercice 1 :")
	// fmt.Println(exercices.Split("I have a cat and a dog"))
	// fmt.Println("Exercice 2 :")
	// fmt.Println(exercices.MoreThanFiveLetters([]string{"dog", "elephant", "cat", "hippopotamus", "giraffe"}))
	fmt.Println("Exercice 3 :")
	fmt.Println("n=0 :", exercices.Fibonacci(0))
	fmt.Println("n=1 :", exercices.Fibonacci(1))
	fmt.Println("n=2 :", exercices.Fibonacci(2))
	fmt.Println("n=100 :", exercices.Fibonacci(10))
	// fmt.Println("Exercice 5 :")
	// exercices.StartWebServer()
	// fmt.Println("Exercice 6 :")
	// exercices.StartAPI()
}
