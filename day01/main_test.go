package main

import "testing"

func TestFindFirstDigit(t *testing.T) {
	actual := findFirstDigit("steve1ok3beep")
	expected := '1'

	if *actual != expected {
		t.Errorf("actual %q expected %q", *actual, expected)
	}
}

func TestFindLastDigit(t *testing.T) {
	actual := findLastDigit("steve1ok3beep")
	expected := '3'

	if *actual != expected {
		t.Errorf("actual %q expected %q", *actual, expected)
	}
}

func TestFindFirstAndLastDigit(t *testing.T) {
	actual := findFirstAndLastDigit("steve1ok53hher53beep")
	expected := int64(13)

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}
