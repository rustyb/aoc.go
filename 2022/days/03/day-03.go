package main

import (
	"fmt"
	// "strings"

	utils "github.com/rustyb/aoc-go-2021/2022"
)


func main() {
	sacks := utils.ReadFile("\n", "./day-03-input.txt")
	sum := 0
	for _, line := range sacks {
		// stringSplit := strings.Split(line, "")
		// get the len of the line
		length := len(line)
		// fmt.Printf("len %d \n", length)
		compartment1, compartment2 := line[:length/2], line[length/2:]
		fmt.Printf("len: %d CPT1: %v CPT2: %v\n", length,  compartment1, compartment2)
		t := findType(compartment1, compartment2)
		fmt.Printf("type: %d \n", t)
		sum += findPriority(t)
	}
	fmt.Printf("Sum %d \n",  sum)
	
	sumPart2 := 0
	for i := 0; i < len(sacks); i += 3 {
		r1, r2, r3 := sacks[i+0], sacks[i+1], sacks[i+2]
		b := findBadge(r1, r2, r3)
		sumPart2 += findPriority(b)
	}
	fmt.Println(sumPart2)
}

func findType(s1, s2 string) byte {
	// TODO:
	// since len(s1)==len(s2) we can just iterate over one string
	for _, a1 := range []byte(s1) {
		for _, a2 := range []byte(s2) {
			if a1 == a2 {
				return a1
			}
		}
	}
	return '.'
}

func findPriority(p byte) int {
	if p >= 97 {
		return (int(p) - 96)
	} else {
		return (int(p) - 38)
	}
}

func findBadge(s1, s2, s3 string) byte {
	for _, a1 := range []byte(s1) {
		for _, a2 := range []byte(s2) {
			for _, a3 := range []byte(s3) {
				if a1 == a2 && a1 == a3 {
					return a1
				}
			}
		}
	}
	return '.'
}
