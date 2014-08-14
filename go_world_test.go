package main

import (
	"testing"
)

func AssertEquals(expect, actual, t *testing.T){
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestDouble(t *testing.T){
	param := 1
	result := Double(param)
	AssertEquals(2, result, t)
}

func TestDouble2(t *testing.T){
	param := 2
	result := Double(param)
	AssertEquals(4, result, t)
}
