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

	for _, elem := range listRight {
		tmp, _ := strconv.Atoi(string(elem[len(elem)-1]))
		if tmp > degree {
			degree = tmp
		}
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

// this function add missing character in each term
func formatEquation(array []string) []string {
	regOnlyDigit := regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]*$`)
	regOnlyX := regexp.MustCompile(`^[-+]?[X]$`)
	regNoPow := regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]\*X*$`)
	regNoCoef := regexp.MustCompile(`^[+-]?X\^[0-9]$`)
	regMissingSign := regexp.MustCompile(`^[0-9]*\.?[0-9]*\*X[\^][0-2]$`)

	for i, elem := range array {
		if regOnlyDigit.MatchString(elem) == true {
			array[i] = elem + "*X^0"
		} else if regOnlyX.MatchString(elem) == true {
			if elem[0] == '-' {
				array[i] = "-1*X^1"
			} else {
				array[i] = "+1*X^1"
			}
		} else if regNoPow.MatchString(elem) == true {
			array[i] = elem + "^1"
		} else if regNoCoef.MatchString(elem) == true {
			if elem[0] == '-' {
				elem = elem[1:len(elem)]
				array[i] = "-1*" + elem
			} else if elem[0] == '+' {
				elem = elem[1:len(elem)]
				array[i] = "+1*" + elem
			} else if elem[0] == 'X' {
				elem = elem[1:len(elem)]
				array[i] = "+1*X" + elem
			}
		} else if regMissingSign.MatchString(elem) == true {
			array[i] = "+" + elem
		}
	}

	return (array)
}

func moveRightToLeft(left []string, right []string) []string {
	for i, elem := range right {
		if strings.Contains(right[i], "+") == true {
			right[i] = strings.Replace(elem, "+", "-", 1)
		} else {
			right[i] = strings.Replace(elem, "-", "+", 1)
		}
	}

	left = append(left, right...)

	return (left)
}

func ReduceForm(equation []string) []string {
	left, right := cleanEquation(equation)

	left = formatEquation(left)
	right = formatEquation(right)

	// move all right side to left side
	left = moveRightToLeft(left, right)

	//sumTerme(left)

	pretty.Println(left)
	pretty.Println(right)

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
