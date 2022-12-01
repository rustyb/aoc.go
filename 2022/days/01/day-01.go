package main

import (
	"fmt"
	"sort"
	"strings"

	utils "github.com/rustyb/aoc-go-2021/2022"
)

func main() {
	lines := utils.ReadFile("\n\n", "./day-01-input.txt")

	// linesAsArrays := strings.Split(lines[0], "\n")
	linesAsArrays := [][]int{}

	for _, ele := range lines { // you can escape index by _ keyword
		stringSplit := strings.Split(ele, "\n")
		asInts := utils.ConvertArrayToInts(stringSplit)

		linesAsArrays = append(linesAsArrays, asInts)
	}

	//  Brute force lets find the elf carrying the most calories

	indexOfGreatestSum := 0
	greatestSum := 0

	for i := 0; i < len(linesAsArrays); i++ {
		fmt.Printf("%v \n", linesAsArrays[i])
		sum := utils.SumValues(linesAsArrays[i])

		if sum >= greatestSum {
			indexOfGreatestSum = i
			greatestSum = sum
		}
	}

	fmt.Printf("Part 1: Elf %d is carrying %d\n", indexOfGreatestSum, greatestSum)

	// part 2 - find the top three sums and sum them up
	sums := []int{}

	for i := 0; i < len(linesAsArrays); i++ {
		fmt.Printf("%v \n", linesAsArrays[i])

		sum := utils.SumValues(linesAsArrays[i])

		sums = append(sums, sum)
	}

	// sortedInts := sort.Reverse(sort.IntSlice(sums[:]))
	sort.Slice(sums, func(a, b int) bool {
		return sums[b] < sums[a]
	})

	fmt.Printf("Top 3 values \n")
	fmt.Printf("%v \n", sums[:3])
	fmt.Printf("%v \n", utils.SumValues(sums[:3]))

}

func ReadFile(s string) {
	panic("unimplemented")
}
