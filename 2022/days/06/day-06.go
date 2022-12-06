package main

import (
	"fmt"

	utils "github.com/rustyb/aoc-go-2021/2022"
)

func main() {
	// EMPTY_NAME := "AAAAAAAAA"
	input := utils.ReadFile("\n", "./day-06-input.txt")

	fmt.Printf("Input array %v\n", input)

	chars := []rune(input[0])

	fmt.Printf("Chars array %v\n", chars)

	result := findMarkerIndex(chars, 4)

	fmt.Printf("Chars array %d\n", result)

	resultB := findMarkerIndex(chars, 14)
	fmt.Printf("Chars array B %d\n", resultB)
}

func Contains[T comparable](inputArray []T, compareInput T) bool {
	for _, v := range inputArray {
		if v == compareInput {
			return true
		}
	}
	return false
}

func findMarkerIndex(chars []rune, markerLength int) int {
	for i := markerLength - 1; i < len(chars); i++ {
		group := chars[i-(markerLength-1) : i+1]
		fmt.Printf("findMarkerIndex group index %d %c\n", i, group)
		if isMarker(group) {
			return i + 1
		}
	}
	return 0
}

func isMarker(group []rune) bool {
	seen := make([]rune, 0, len(group))
	for _, r := range group {
		if Contains(seen, r) {
			return false
		}
		seen = append(seen, r)
	}
	return true
}
