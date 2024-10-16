package main

import (
	"fmt"
	"math"
	"maxscore/heap"
)

func maxKelements(nums []int, k int) int {
	h := heap.NewHeap()
	score := 0

	for _, num := range nums {
		h.Add(num)
	}

	for k > 0 {
		value, ok := h.Delete()
		if !ok {
			break
		}

		score += value
		newValue := int(math.Ceil(float64(value) / 3))
		h.Add(newValue)
		k--
	}

	return score

}

func main() {
	list := []int{10, 10, 10, 10, 10}
	k := 5
	score := maxKelements(list, k)
	fmt.Printf("Max score: %d\n", score)
}
