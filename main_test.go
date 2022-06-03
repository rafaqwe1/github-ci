package main

import "testing"

func TestSum(t *testing.T) {
	expectedResult := 11
	result := Sum(2, 8)
	if expectedResult != result {
		t.Fatalf(`Error on validation of Sum function, expected %d and received %d`, expectedResult, result)
	}
}
