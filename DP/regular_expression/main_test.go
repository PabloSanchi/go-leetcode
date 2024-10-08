package main

import "testing"

func TestFailure(t *testing.T) {
	s := "aa"
	p := "a"

	if isMatch(s, p) {
		t.Errorf("Expected false, got true")
	}
}

func TestSuccessCase1(t *testing.T) {
	s := "aa"
	p := "a*"

	if !isMatch(s, p) {
		t.Errorf("Expected true, got false")
	}
}

func TestSuccessCase2(t *testing.T) {
	s := "ab"
	p := ".*"

	if !isMatch(s, p) {
		t.Errorf("Expected true, got false")
	}
}
