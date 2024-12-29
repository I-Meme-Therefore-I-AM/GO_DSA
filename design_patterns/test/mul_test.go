package main

import "testing"


func TestMul(t *testing.T) {
	a := 5
	b := 6

	expected := 30

	res := multiply(a, b)

	if res != expected{
		t.Errorf("Our multiply function doesn't work, %d*%d isn't %d\n", a, b, res)
	}
}
