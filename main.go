package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kr/pretty"
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

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func cleanEquation(equation []string) ([]string, []string) {
	// remove blank
	left := removeSpace(equation[0])
	right := removeSpace(equation[1])

	// add space to keep sign
	left = addSpaceBeforeSign(left)
	right = addSpaceBeforeSign(right)

	// split string on space
	listLeft := strings.Split(left, " ")
	listRight := strings.Split(right, " ")

	// Remove blank slice
	if listLeft[0] == "" {
		listLeft = remove(listLeft, 0)
	}
	if listRight[0] == "" {
		listRight = remove(listRight, 0)
	}

	return listLeft, listRight
}

func formatEquation(left []string, right []string) string {
	regOnlyDigit := regexp.MustCompile(`[-+]?[0-9]*\.?[0-9]`)

	for i, elem := range left {
		if regOnlyDigit.Match([]byte(elem)) == true {
			left[i] = elem + "*X^0"
		}
	}

	//pretty.Println(left)
	return ""
}

func ReduceForm(equation []string) []string {
	left, right := cleanEquation(equation)

	pretty.Println(left)
	formatEquation(left, right)

	//if len(left) > 3 || len(right) > 3 {
	//log.Printf("Equation bad formatted\n")
	//return ""
	//}

	degree := getDegree(left, right)
	if degree > 2 {
		fmt.Println("Polynomial degree: 3\nThe polynomial degree is stricly greater than 2, I can't solve")
		return nil
	}

	//pretty.Println(right)

	return left
}

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

	reduced := ReduceForm(split)
	if reduced == nil {
		return
	}
}

// match good formated term
// [-+]?[0-9]*\.?[0-9]*\*X[\^][0-2]
