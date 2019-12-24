package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func removeSpace(str string) string {
	reg := regexp.MustCompile(`[ ]`)
	newStr := reg.ReplaceAllString(str, "")

	return (newStr)
}

func addSpaceBeforeSign(str string) string {
	// add space before +
	reg := regexp.MustCompile(`[+]`)
	str = reg.ReplaceAllLiteralString(str, " +")

	// add space before -
	reg = regexp.MustCompile(`[-]`)
	str = reg.ReplaceAllLiteralString(str, " -")

	return (str)
}

func getDegree(listLeft []string, listRight []string) int {
	var degree int

	for _, elem := range listLeft {
		tmp, _ := strconv.Atoi(string(elem[len(elem)-1]))
		if tmp > degree {
			degree = tmp
		}
	}

	// TODO : fix crash here
	for _, elem := range listRight {
		tmp, _ := strconv.Atoi(string(elem[len(elem)-1]))
		if tmp > degree {
			degree = tmp
		}
		// addition chaque terme
		//add := strings.Split(elem, "*")
		//fmt.Println(add)
		//number, _ := strconv.Atoi(add[0])
		//fmt.Println(number)
	}

	return (degree)
}

func reduceForm(equation []string) string {
	// add space to keep sign
	left := addSpaceBeforeSign(equation[0])
	right := addSpaceBeforeSign(equation[1])

	// remove blank
	left = removeSpace(left)
	right = removeSpace(right)

	// remove space
	listLeft := strings.Split(left, " ")
	listRight := strings.Split(right, " ")

	fmt.Println(listLeft)
	fmt.Println(listRight)

	if len(listLeft) > 3 || len(listRight) > 3 {
		fmt.Fprintf(os.Stderr, "Equation bad formatted\n")
		return ""
	}

	degree := getDegree(listLeft, listRight)
	if degree > 2 {
		fmt.Println("Polynomial degree: 3\nThe polynomial degree is stricly greater than 2, I can't solve")
		return ""
	}

	//for elem := range listRight {

	//}

	return "a"
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

	reduced := reduceForm(split)
	if reduced == "" {
		return
	}
}
