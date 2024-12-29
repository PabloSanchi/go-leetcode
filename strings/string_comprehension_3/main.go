package main

import (
	"fmt"
	"strings"
)

func compressedString(word string) string {
	var comp strings.Builder
	counter := 1
	char := rune(word[0])

	for _, c := range word[1:] {
		if c == char {
			if counter < 9 {
				counter++
			} else {
				comp.WriteByte(byte(counter) + '0')
				comp.WriteRune(char)
				counter = 1
			}
		} else {
			comp.WriteByte(byte(counter) + '0')
			comp.WriteRune(char)
			counter = 1
			char = c
		}
	}

	comp.WriteByte(byte(counter) + '0')
	comp.WriteRune(char)

	return comp.String()
}

func main() {
	input := "abcde"
	fmt.Println(compressedString(input))
}
