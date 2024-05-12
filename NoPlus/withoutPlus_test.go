package main

import "testing"

func TestThatFunctionReturnSumWithoutPlus(t *testing.T) {
	expected := 10
	got := getSum(5, 5)

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestAddFunctionOnNegative(t *testing.T) {
	expected := 0
	got := getSum(-5, 5)

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
