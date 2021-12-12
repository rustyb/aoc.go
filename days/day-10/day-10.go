package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-10-input.txt"
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

func main() {

	inputs := ReadFile("\n")
	score := 0
	for _, line := range inputs {
		stack := []rune{}
		corrupted := rune(0)

	loop:
		for _, c := range line {
			switch c {
			case '(', '[', '{', '<':
				stack = append(stack, c)
			case ')', ']', '}', '>':
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if c == ')' && pop == '(' ||
					c == ']' && pop == '[' ||
					c == '}' && pop == '{' ||
					c == '>' && pop == '<' {
					continue
				} else {
					corrupted = c
					break loop
				}

			}

		}
		fmt.Println(line)
		switch corrupted {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		default:
		}
	}

	fmt.Println("Whoop whoop a score of ", score)
}
