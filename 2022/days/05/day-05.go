package main

import (
	"fmt"
	"strings"

	utils "github.com/rustyb/aoc-go-2021/2022"
)

// func isOverlapping(inputA []int, inputB []int) bool {
// 	min := inputA[0]
// 	max := inputA[1]

// 	return inputB[0] >= min && inputB[1] <= max || min >= inputB[0] && max <= inputB[1]

// }

// func isPartiallyOverlapping(inputA []int, inputB []int) bool {
// 	min := inputA[0]
// 	max := inputA[1]

// 	return inputB[0] <= max && inputB[1] >= min || min <= inputB[1] && max >= inputB[1]

// }

// N is an alias for an unallocated struct
func N(size int) []struct{} {
	return make([]struct{}, size)
}

// function, which takes a string as
// argument and return the reverse of string.
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func push(s []string, r string) {
	s = append(s, r)
}

func pop() (s []string, r string) {
	r = s[len(s)-1]
	s = s[:len(s)-1]
	return
}

func makeBoard(cratesNames []string) map[int][]string {
	board := make(map[int][]string)

	for index, nameArray := range cratesNames {
		strRev := reverse(nameArray)
		fmt.Printf("cratesNames Rev %s \n", strRev[1:])
		nameArray := strings.Split(strRev[1:], ",")

		crateIndex := 9

		fmt.Printf("%d: cratesNames %v %d\n", index, nameArray, len(nameArray))

		for _, name := range nameArray {
			if name != "AAAA" {
				fmt.Println(name)
				board[crateIndex] = append(board[crateIndex], name)
			}
			crateIndex -= 1
		}

	}

	return board
}

func main() {
	// EMPTY_NAME := "AAAAAAAAA"
	input := utils.ReadFile("\n\n", "./day-05-input.txt")

	inputMap := input[0]
	actions := input[1]

	// fmt.Printf("InputMap len %s\n", inputMap)
	// fmt.Printf("InputMap %s \n", actions)
	// withoutSpaces := strings.ReplaceAll(inputMap, " ", "")
	withoutSpaces := strings.ReplaceAll(inputMap, "    ", "[AAAA] ")
	fmt.Printf("InputMap with AAA \n%s\n", withoutSpaces)
	withoutSpaces = strings.ReplaceAll(withoutSpaces, " ", "")
	withoutLeftBrackets := strings.ReplaceAll(withoutSpaces, "[", "")
	withoutRightBrackets := strings.ReplaceAll(withoutLeftBrackets, "]", ",")
	fmt.Printf("InputMap withoutRightBrackets \n%s\n", withoutRightBrackets)

	mapSplit := strings.Split(withoutRightBrackets, "\n")
	fmt.Printf("mapsplit %v \n", mapSplit)

	cratesIndex := utils.ConvertArrayToInts(strings.Split(mapSplit[len(mapSplit)-1], "AAAA,"))
	fmt.Printf("crates %v leng %d \n", cratesIndex, len(cratesIndex))

	cratesNames := mapSplit[:len(mapSplit)-1]
	fmt.Printf("cratesNames %v \n", cratesNames)

	// board := make(map[int][]string)

	// for index, nameArray := range cratesNames {
	// 	strRev := reverse(nameArray)
	// 	fmt.Printf("cratesNames Rev %s \n", strRev[1:])
	// 	nameArray := strings.Split(strRev[1:], ",")

	// 	crateIndex := 9

	// 	fmt.Printf("%d: cratesNames %v %d\n", index, nameArray, len(nameArray))

	// 	for _, name := range nameArray {
	// 		if name != "AAAA" {
	// 			fmt.Println(name)
	// 			board[crateIndex] = append(board[crateIndex], name)
	// 		}
	// 		crateIndex -= 1
	// 	}

	// }
	board := makeBoard(cratesNames)
	fmt.Printf("board Rev %v \n", board)

	fmt.Printf("InputMap len %v\n")

	splitActions := strings.Split(actions, "\n")

	for _, action1 := range splitActions {
		item := strings.Split(action1, " ")
		fmt.Printf("Item %v len: %d\n", item, len(item))

		act, toMove, from, to := item[0], utils.ConvertStringToInt(item[1]), utils.ConvertStringToInt(item[3]), utils.ConvertStringToInt(item[5])
		// to := item[5]
		fmt.Printf("action: %s  move: %d from: %d to: %d\n", act, toMove, from, to)

		// move them one by one
		//Move elements one by one
		for move := 0; move < toMove; move++ {
			// board[to-1]
			// fmt.Printf("Will move to: %v\n", board[to])
			// board[to] = append(board[to], board[from][0])
			board[to] = append([]string{board[from][0]}, board[to]...)
			board[from] = board[from][1:]
			// stacks[to-1].push(stacks[from-1].pop())
			// fmt.Printf("After move to: %v\n", board[to])
			// fmt.Printf("After move to: %v\n", board[from])
		}
		fmt.Printf("AFTER BOARD: %v\n", board)

	}
	fmt.Printf("FINAL BOARD: %v\n", board)

	// Part 2
	board = makeBoard(cratesNames)

	splitActions1 := strings.Split(actions, "\n")

	for _, action2 := range splitActions1 {
		item := strings.Split(action2, " ")
		fmt.Printf("Item %v len: %d\n", item, len(item))

		act, toMove, from, to := item[0], utils.ConvertStringToInt(item[1]), utils.ConvertStringToInt(item[3]), utils.ConvertStringToInt(item[5])

		fmt.Printf("----PART 2 action: %s  move: %d from: %d to: %d\n", act, toMove, from, to)
		fmt.Printf("INPUT BOARD: %v\n", board)

		endIndex := toMove

		fmt.Printf("FROM %d: %v\n", from, board[from][:endIndex])
		fmt.Printf("FROM Remain %d: %v\n", from, board[from][endIndex:])
		fmt.Printf("To %d: %v\n", to, board[to])

		fromBefore := board[from][:endIndex]
		fromAfter := board[from][endIndex:]
		c := make([]string, len(fromAfter))
		copy(c, fromAfter)

		board[to] = append(fromBefore, board[to]...)
		fmt.Printf("To After %d: %v\n", to, board[to])
		board[from] = c
		fmt.Printf("FROM After Remain %d: %v\n", from, board[from])

	}
	fmt.Printf("FINAL BOARD 2: %v\n", board)
}
