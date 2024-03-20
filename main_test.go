package main

import "testing"

func TestSum(t *testing.T) {
	test := sum(20, 30)
	result := 50.0
	if test != result {
		t.Error("Expected value ", result, ", got value ", test)
	}
}
func TestSubtract(t *testing.T) {
	test := subtract(20, 30)
	result := -10.0
	if test != result {
		t.Error("Expected value ", result, ", got value ", test)
	}
}
func TestMultiply(t *testing.T) {
	test := multiply(20, 30)
	result := 600.0
	if test != result {
		t.Error("Expected value ", result, ", got value ", test)
	}
}
func TestDivide(t *testing.T) {
	test, _ := divide(60, 20)
	result := 3.0
	if test != result {
		t.Error("Expected value ", result, ", got value ", test)
	}
}
