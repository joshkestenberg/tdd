package main

import (
	"testing"
)

type example struct {
	expression string
	result     int
}

var examples = []example{
	example{"(+ 1 5)", 6},
	example{"(+ (- 9 7) 8)", 10},
	example{"(- (* (+ 2 2) (+ 2 3)) (% 30 3)", 10},
}

func TestCalc(t *testing.T) {
	for _, e := range examples {
		expr := e.expression
		result := e.result
		result_ := Calculate(expr)
		if result != result_ {
			t.Fatalf("Got %v; Expected %v", result_, result)
		}
	}
}
