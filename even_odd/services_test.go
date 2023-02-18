package main

import "testing"

func TestNewElement(t *testing.T) {
	el := newElement()
	
	if len(el) != 11 {
		t.Errorf("Expect len of el to be 11 but got %d", len(el))
	}

	if el[0] != 0 {
		t.Errorf("Expect first el to be 0 but got %d", el[0])
	}

	if el[len(el) - 1] != 10 {
		t.Errorf("Expect last el to be 0 but got %d", el[len(el) - 1])
	}
}

func TestPrintNumberType(t *testing.T) {
	evenNumber := printNumberType(10)
	oddNumber := printNumberType(1)

	if evenNumber != "10 is even" {
		t.Errorf("Invalid result got %s", evenNumber)
	}

	if oddNumber != "1 is odd" {
		t.Errorf("Invalid result got %s", oddNumber)
	}
}
