package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage : %s 5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0\n", os.Args[0])
		return
	}
	equation := os.Args[1]

	split := strings.Split(equation, "= ")

	if len(split) != 2 {
		log.Printf("Equation bad formatted\n")
		return
	}

	A, B, C, err := ReduceForm(split)
	if err != nil {
		return
	}

	fmt.Printf("A = %v, B = %v, C = %v\n", A, B, C)
	fmt.Println("Formula to get delta is : b² - 4ac")

	delta := getDelta(A, B, C)
	_ = delta
}
