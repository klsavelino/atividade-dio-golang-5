package main

import (
	"errors"
	"fmt"
	"log"
)

func sum(x float64, y float64) float64 {
	return x + y
}
func subtract(x float64, y float64) float64 {
	return x - y
}
func multiply(x float64, y float64) float64 {
	return x * y
}
func divide(x float64, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return x / y, nil
}

func main() {
	var operation int
	var x float64
	var y float64
	var result float64
	var err error

	fmt.Print("Enter the first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&y)

	fmt.Print("1. Sum\n2.Subtract\n3.Multiply\n4.Divide\n")
	fmt.Print("Choose the operation: ")
	fmt.Scan(&operation)

	switch operation {
	case (1):
		result = sum(x, y)
	case (2):
		result = subtract(x, y)
	case (3):
		result = multiply(x, y)
	case (4):
		result, err = divide(x, y)
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("No valid option was chosen, try again or exit the program")
	}
	fmt.Printf("The result is %.2f", result)
	fmt.Print("\n")
}
