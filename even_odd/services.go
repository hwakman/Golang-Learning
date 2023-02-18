package main

import "fmt"


type elements []int

func newElement() elements {
	var newElement elements
	for i := range [11]int{} {
		newElement = append(newElement, i)
	}
	return newElement
}

func printNumberType(num int) string {
	if num % 2 == 0 {
		return fmt.Sprintf("%d is even", num)
	}
	return fmt.Sprintf("%d is odd", num)
}

func (e elements) printEvenAndOddElements() []error {
	var errors []error
	for _, el := range e {
		_, err := fmt.Println(printNumberType(el))
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
