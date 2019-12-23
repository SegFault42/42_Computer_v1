package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func reduceForm(equation []string) string {
	// get a b and c
	reg := regexp.MustCompile(`[+-]`)
	trinome := reg.Split(equtrinometion[0], -1)
	for i := rtrinomenge trinome {
		reg := regexp.MustCompile(`[ ]`)
		trinome[i] = reg.RepltrinomeceAllString(trinome[i], "")
		fmt.Println(trinome[i])
	}

	return ""
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage : %s 5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0\n", os.Args[0])
		return
	}
	equation := os.Args[1]

	split := strings.Split(equation, "= ")

	if len(split) != 2 {
		fmt.Fprintf(os.Stderr, "Equation bad formatted\n")
		return
	}

	reduceForm(split)
}
