package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func isEquPossible(array []string) bool {
	var result float32

	for _, elem := range array {
		i, _ := strconv.ParseFloat(elem, 64)
		result += float32(i)
	}

	if result != 0 {
		return (false)
	}

	return (true)
}

func getCoefAndPow(array []string) ([]string, []string, []string, error) {
	var powZero []string
	var powOne []string
	var powTwo []string

	reg := regexp.MustCompile(`[-+]?[0-9]*\.?[0-9]*`)
	err := errors.New("There are no solution")

	// get coef and pow
	for _, elem := range array {
		num := reg.FindAllString(elem, -1)
		num = removeBlankFromSlice(num)
		if len(num) != 2 {
			log.Printf("%v : Parsing error\n", num)
			return nil, nil, nil, err
		}
		if num[len(num)-1] == "0" {
			powZero = append(powZero, num[0])
		} else if num[len(num)-1] == "1" {
			powOne = append(powOne, num[0])
		} else if num[len(num)-1] == "2" {
			powTwo = append(powTwo, num[0])
		} else {
			log.Printf("Polynomial degree: %s. The polynomial degree is stricly greater than 2, I can't solve\n", num[len(num)-1])
			return nil, nil, nil, err
		}
	}

	// check if equ degree 0 solvable
	if len(powOne) == 0 && len(powTwo) == 0 {
		if isEquPossible(powZero) == false {
			err := errors.New("There are no solution")
			return nil, nil, nil, err
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

	return powZero, powOne, powTwo, nil
}

func sumTerm(array []string) (float32, float32, float32, error) {
	var A, B, C float32

	powZero, powOne, powTwo, err := getCoefAndPow(array)
	if err != nil {
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

	if A == 0 {
		err := errors.New("A cannot be 0")
		return 0, 0, 0, err
	}

	return A, B, C, nil
}

func getDelta(A, B, C float32) float32 {
	fmt.Printf("Get delta : (%v²) - (4 * (%v) * (%v))\n", B, A, C)

	delta := (B * B) - (4 * (A) * (C))

	fmt.Printf("delta is %v\n", delta)

	return (delta)
}

func sqrt(x float32) float32 {
	z := 2 - (2*2-x)/(2*2)
	for zn, delta := z, z; delta > 0.00001; z = zn {
		zn = z - (z*z-x)/(2*z)
		delta = z - zn
	}
	return (z)
}

func resultPositiveDelta(A, B, delta float32) {
	fmt.Println("Delta is greater than 0.\nTwo results are possible.")
	fmt.Println("Formula to get result is : -b -√Δ / 2a")
	fmt.Printf("Get result : %v - √%v / 2 * %v\n", (-1 * B), sqrt(delta), A)
	fmt.Println("Result is :")
	fmt.Println(((-1 * B) - sqrt(delta)) / (2 * A))

	fmt.Println("Or :\nFormula to get result is : -b +√Δ / 2a")
	fmt.Printf("Get result : %v + √%v / 2 * %v\n", (-1 * B), sqrt(delta), A)
	fmt.Println("Result is :")
	fmt.Println(((-1 * B) + sqrt(delta)) / (2 * A))
}

func resultZeroDelta(A, B float32) {
	fmt.Println("Delta equal 0.")
	fmt.Println("Formula to get result is : -b / 2a")
	fmt.Printf("Get result : %v / 2 * %v\n", -1*B, A)
	fmt.Println("Result is :")
	fmt.Println((-1 * B) / (2 * A))
}

func resultNegativeDelta(A, B, delta float32) {
	fmt.Println("Delta is lower than 0.\nTwo results are possible.")
	fmt.Println("Formula to get result is : -b + i√Δ / 2a")

	fmt.Println("Result is :")
	//fmt.Printf("%v + i√%v / %v\n", -1*B, delta, 2*A)
	fmt.Printf("%v + %v * i\n", -B/(2*A), sqrt(-delta)/(2*A))
	fmt.Printf("Or:\n%v - %v * i\n", -B/(2*A), sqrt(-delta)/(2*A))

}

func calculateResult(A, B, delta float32) {
	if delta < 0 {
		fmt.Println("Delta is inferior than 0, no solutions for this equation")
		resultNegativeDelta(A, B, delta)
	} else if delta > 0 {
		resultPositiveDelta(A, B, delta)
	} else {
		resultZeroDelta(A, B)
	}
}
