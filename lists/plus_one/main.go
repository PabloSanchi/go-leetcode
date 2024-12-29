package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	carry := 1

	index := len(digits) - 1

	for index >= 0 {
		value := digits[index]
		if value == 9 {
			carry = 1
			digits[index] = 0
		} else {
			digits[index] += carry
			return digits
		}

		index--
	}

	digits = append([]int{1}, digits...)

	return digits
}

func main() {
	output := plusOne([]int{2, 4, 9, 3, 9})
	fmt.Printf("Output: %v", output)
}
