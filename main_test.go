package main

import (
	"regexp"
	"testing"
)

type test struct {
	input    []string
	expected bool
}

func TestReducedForm(t *testing.T) {
	var tests = []test{
		{{"-5 * X^0"}, true},
		{{"+4 * X^1"}, true},
		{{"+90"}, false},
	}

	reg := regexp.MustCompile(`[-+]?[0-9]*\.?[0-9]*\*X[\^][0-2]`)

	for _, elem := range tests {
		reduced := ReduceForm(elem.input)
		for _, item := range reduced {
			if reg.Match([]byte(item)) == false {
				t.Error("Equation Bad formated")
			}
		}
	}
}
