package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/kr/pretty"
)

func TestReducedForm(t *testing.T) {
	var tests = []string{
		"-5 * X^0",
		"+4 * X^1",
		"+90",
	}

	reg := regexp.MustCompile(`[-+]?[0-9]*\.?[0-9]*\*X[\^][0-2]`)

	reduced := ReduceForm(tests)
	pretty.Println("red =", reduced)
	for _, item := range reduced {
		fmt.Println(item)
		if reg.Match([]byte(item)) == false {
			t.Error("Equation Bad formated")
		}
	}
}
