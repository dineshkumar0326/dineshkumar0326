package main

import "testing"

func TestHello(t *testing.T) {
	emptyResult := hello("")

	if emptyResult != "Hello " {
		t.Errorf("Test failed, expected %v, got %v", "Hello", emptyResult)
	}

	result := hello("Mike")

	if result != "Hello Mike" {
		t.Errorf("Test failed, expected %v, got %v", "Hello Mike", result)
	}
}
