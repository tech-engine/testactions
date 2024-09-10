package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	// this is a comment
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
