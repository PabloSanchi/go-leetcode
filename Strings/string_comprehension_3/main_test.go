package main

import "testing"

func TestCase1(t *testing.T) {
	input := "abcde"
	expected := "1a1b1c1d1e"
	result := compressedString(input)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestCase2(t *testing.T) {
	input := "aaaaaaaaaaaaaabb"
	expected := "9a5a2b"
	result := compressedString(input)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
