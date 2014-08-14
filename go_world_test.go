package main

import (
	"testing"
)

func TestDouble(t *testing.T){
	param := 1
	result := Double(param)

	if result != 2 {
		t.Errorf("Double(%d) = %d, want %d", param, result, 2)
	}
}

func TestDouble2(t *testing.T){
	param := 2
	result := Double(param)

	if result != 4 {
		t.Errorf("Double(%d) = %d, want %d", param, result, 4)
	}
}
