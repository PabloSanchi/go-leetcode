package main

import "testing"

func TestCase1(t *testing.T) {
	list := []int{10, 10, 10, 10, 10}
	k := 5
	score := maxKelements(list, k)
	if score != 50 {
		t.Errorf("Expected 50, but got %d", score)
	}
}

func TestCase2(t *testing.T) {
	list := []int{1, 10, 3, 3, 3}
	k := 3
	score := maxKelements(list, k)
	if score != 17 {
		t.Errorf("Expected 17, but got %d", score)
	}
}
