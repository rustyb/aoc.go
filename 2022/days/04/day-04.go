package main

import (
	"fmt"
	"strings"

	utils "github.com/rustyb/aoc-go-2021/2022"
)

func isOverlapping(inputA []int, inputB []int) bool {
	min := inputA[0]
	max := inputA[1]

	return inputB[0] >= min && inputB[1] <= max || min >= inputB[0] && max <= inputB[1]

}

func isPartiallyOverlapping(inputA []int, inputB []int) bool {
	min := inputA[0]
	max := inputA[1]

	return inputB[0] <= max && inputB[1] >= min || min <= inputB[1] && max >= inputB[1]

}

func main() {
	assignments := utils.ReadFile("\n", "./day-04-input.txt")

	overlappingCount := 0
	partialOverlappingCount := 0

	for _, assignment := range assignments {
		assignmentSplit := strings.Split(assignment, ",")
		fmt.Printf("Split: %v \n", assignmentSplit)

		assignmentA := strings.Split(assignmentSplit[0], "-")
		assignmentB := strings.Split(assignmentSplit[1], "-")

		assignmentAInt := utils.ConvertArrayToInts(assignmentA)
		assignmentBInt := utils.ConvertArrayToInts(assignmentB)

		if isOverlapping(assignmentAInt, assignmentBInt) == true {
			overlappingCount += 1
		}

		if isPartiallyOverlapping(assignmentAInt, assignmentBInt) == true {
			partialOverlappingCount += 1
		}

	}
	fmt.Printf("Overlapping %d \n", overlappingCount)
	fmt.Printf("Partial Overlapping %d \n", partialOverlappingCount)

}
