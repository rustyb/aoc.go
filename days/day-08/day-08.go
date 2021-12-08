package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-08-input.txt"
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	// fmt.Println(fileContent)
	stringSplit := strings.Split(fileContent, delimeter)

	// To deal with the \n on the end of the file read in
	fmt.Printf("%q\n", stringSplit[:len(stringSplit)-1])
	return stringSplit[:len(stringSplit)-1]
}

func countDigits(digits []string) int {
	var counts = make([]int, 9)
	for _, v := range digits {
		if len(v) == 7 {
			counts[8]++
		}
		if len(v) == 2 {
			counts[1]++
		}
		if len(v) == 4 {
			counts[4]++
		}
		if len(v) == 3 {
			counts[7]++
		}

	}

	var sum = 0
	for _, i := range counts {
		sum += i
	}

	return sum
}

func main() {
	inputs := ReadFile("\n")
	sumTotal := 0
	for _, v := range inputs {
		parts := strings.Split(v, " | ")
		digits := strings.Split(parts[1], " ")
		counts := countDigits(digits)
		sumTotal += counts
		fmt.Println(digits)
		fmt.Println(counts)
	}
	fmt.Println("Total number of 1,4,7,8 is", sumTotal)

	fmt.Println("part2 (input):", part2())
}

// I cannot take credit for the second part
// Code is sourced from https://github.com/alextanhongpin/advent-of-code-2021/blob/main/day08/main.go

func part2() int {
	lines := ReadFile("\n")
	var result int
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		in := strings.Fields(strings.TrimSpace(parts[0]))
		out := strings.Fields(strings.TrimSpace(parts[1]))

		segmentsByNumber := make(map[int]string)
		for len(segmentsByNumber) != 10 {
			for _, n := range in {
				if len(n) == 2 {
					segmentsByNumber[1] = n
				}
				if len(n) == 4 {
					segmentsByNumber[4] = n
				}
				if len(n) == 3 {
					segmentsByNumber[7] = n
				}
				if len(n) == 7 {
					segmentsByNumber[8] = n
				}
				// 2, 3 or 5
				if len(n) == 5 {
					one, ok := segmentsByNumber[1]
					if !ok {
						continue
					}
					if len(Set(one).Intersect(Set(n))) == 2 {
						segmentsByNumber[3] = n
						continue
					}
					four, ok := segmentsByNumber[4]
					if !ok {
						continue
					}
					if len(Set(four).Intersect(Set(n))) == 2 {
						segmentsByNumber[2] = n
					} else {
						segmentsByNumber[5] = n
					}
				}
				if len(n) == 6 {
					four, ok := segmentsByNumber[4]
					if !ok {
						continue
					}
					if len(Set(four).Intersect(Set(n))) == 4 {
						segmentsByNumber[9] = n
						continue
					}
					one, ok := segmentsByNumber[1]
					if !ok {
						continue
					}
					if len(Set(one).Intersect(Set(n))) == 2 {
						segmentsByNumber[0] = n
					} else {
						segmentsByNumber[6] = n
					}
				}
			}
		}
		fmt.Println("segmentsBynumber", segmentsByNumber)
		numberBySegments := make(map[string]int)
		for n, seg := range segmentsByNumber {
			numberBySegments[SortedString(seg).Sort()] = n
		}
		var total int
		for _, seg := range out {
			total *= 10
			total += numberBySegments[SortedString(seg).Sort()]
		}
		result += total
	}
	return result
}

type SortedString string

func (s SortedString) Sort() string {
	chars := strings.Split(string(s), "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

type Set string

func (s Set) toMap() map[rune]bool {
	result := make(map[rune]bool)
	for _, c := range s {
		result[c] = true
	}
	return result
}

func (s Set) Difference(other Set) string {
	var result string
	found := make(map[rune]bool)
	tgt := other.toMap()
	for c := range s.toMap() {
		if found[c] {
			continue
		}
		if tgt[c] {
			continue
		}
		result += string(c)
		found[c] = true
	}
	return result
}

func (s Set) Intersect(other Set) string {
	if len(other) > len(s) {
		return other.Intersect(s)
	}
	var result string
	found := make(map[rune]bool)
	tgt := other.toMap()
	for c := range s.toMap() {
		if found[c] {
			continue
		}
		if tgt[c] {
			result += string(c)
			found[c] = true
		}
	}
	return result
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
