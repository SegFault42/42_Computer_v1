package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func getCoefAndPow(array []string) ([]string, []string, []string) {
	var powZero []string
	var powOne []string
	var powTwo []string

	reg := regexp.MustCompile(`[-+]?[0-9]*\.?[0-9]*`)

	// get coef and pow
	for _, elem := range array {
		num := reg.FindAllString(elem, -1)
		num = removeBlankFromSlice(num)
		if len(num) != 2 {
			log.Printf("%v : Parsing error\n", num)
			return nil, nil, nil
		}
		if num[len(num)-1] == "0" {
			powZero = append(powZero, num[0])
		} else if num[len(num)-1] == "1" {
			powOne = append(powOne, num[0])
		} else if num[len(num)-1] == "2" {
			powTwo = append(powTwo, num[0])
		} else {
			log.Printf("Polynomial degree: %s. The polynomial degree is stricly greater than 2, I can't solve\n", num[len(num)-1])
			return nil, nil, nil
		}
	}

	// Replace by one if missing
	if len(powTwo) == 0 {
		powTwo = append(powTwo, "1")
	}
	if len(powOne) == 0 {
		powOne = append(powOne, "1")
	}
	if len(powZero) == 0 {
		powZero = append(powZero, "1")
	}

	return powZero, powOne, powTwo
}

func sumTerm(array []string) (float32, float32, float32, error) {
	var A, B, C float32

	powZero, powOne, powTwo := getCoefAndPow(array)
	if powZero == nil {
		err := errors.New("sumTerm() error")
		return 0, 0, 0, err
	}

	for _, elem := range powZero {
		i, _ := strconv.ParseFloat(elem, 64)
		C += float32(i)
	}
	for _, elem := range powOne {
		i, _ := strconv.ParseFloat(elem, 64)
		B += float32(i)
	}
	for _, elem := range powTwo {
		i, _ := strconv.ParseFloat(elem, 64)
		A += float32(i)
	}

	return A, B, C, nil
}

func getDelta(A, B, C float32) float32 {
	fmt.Printf("Get delta : (%v²) - (4 * (%v) * (%v))\n", B, A, C)

	delta := (B * B) - (4 * (A) * (C))

	fmt.Printf("delta is %v\n", delta)

	return (delta)
}
