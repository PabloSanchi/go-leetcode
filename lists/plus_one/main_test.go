package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 2, 3}
	expectedOutput := []int{1, 2, 4}
	output := plusOne(input)

	if !reflect.DeepEqual(expectedOutput, output) {
		t.Errorf("The expected output is %v, got %v", expectedOutput, output)
	}

}

func TestCase2(t *testing.T) {
	input := []int{4, 3, 2, 1}
	expectedOutput := []int{4, 3, 2, 2}
	output := plusOne(input)

	if !reflect.DeepEqual(expectedOutput, output) {
		t.Errorf("The expected output is %v, got %v", expectedOutput, output)
	}

}

func TestCase3(t *testing.T) {
	input := []int{9}
	expectedOutput := []int{1, 0}
	output := plusOne(input)

	if !reflect.DeepEqual(expectedOutput, output) {
		t.Errorf("The expected output is %v, got %v", expectedOutput, output)
	}

}
