package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(5, 6)

	if result != 11 {
		t.Errorf("Test failed, expected %v, got %v", 11, result)
	}
}
