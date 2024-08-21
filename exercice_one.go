package main

import (
	"fmt"
	"strings"
)

func split(x string) []string {
	return strings.Split(x, " ")
}

func main() {
	fmt.Println(split("I have a cat and a dog"))
}
