package main

import (
	"errors"
	"fmt"
	"log"
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

func removeBlankFromSlice(s []string) []string {
	var array []string

	for _, elem := range s {
		if elem != "" {
			array = append(array, elem)
			//s = append(s[:i], s[i+1:]...)
		}
	}

	return (array)
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

func checkFormatEquation(array []string) bool {
	reg := regexp.MustCompile(`[-+]?[0-9]*\.?[0-9]*\*X[\^][0-9]`)

	for _, elem := range array {
		if reg.MatchString(elem) == false {
			log.Printf("%s bad formatted\n", elem)
			return (false)
		}
	}

	return (true)
}

func printReducedForm(A float32, B float32, C float32) {
	fmt.Printf("Reduced form: %v * X^2 ", A)

	if B > 0 {
		fmt.Printf("+ %v * X^1", B)
	} else {
		fmt.Printf("- %v * X^1", B*-1)
	}

	if C > 0 {
		fmt.Printf("+ %v * X^0 = 0\n", C)
	} else {
		fmt.Printf("- %v * X^0 = 0\n", C*-1)
	}

	fmt.Printf("Reduced form: %vx² ", A)

	if B > 0 {
		fmt.Printf("+ %vx ", B)
	} else {
		fmt.Printf("- %vx", B*-1)
	}

	if C > 0 {
		fmt.Printf("+ %v = 0\n\n", C)
	} else {
		fmt.Printf("- %v = 0\n\n", C*-1)
	}
}

func ReduceForm(equation []string) (float32, float32, float32, error) {
	left, right := cleanEquation(equation)
	err := errors.New("")

	left = formatEquation(left)
	right = formatEquation(right)

	// move all right side to left side
	left = moveRightToLeft(left, right)

	if checkFormatEquation(left) == false {
		return 0, 0, 0, err
	}

	A, B, C, err := sumTerm(left)
	if err != nil {
		return 0, 0, 0, err
	}

	printReducedForm(A, B, C)

	return A, B, C, nil
}
