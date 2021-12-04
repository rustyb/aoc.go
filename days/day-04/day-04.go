package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-04-input.txt"
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	// fmt.Println(fileContent)
	stringSplit := strings.Split(fileContent, delimeter)

	// To deal with the \n on the end of the file read in
	// fmt.Printf("%q\n", stringSplit[:len(stringSplit)-1])
	return stringSplit[:len(stringSplit)-1]
}

// convert the string array to ints to be able to
func convertArrayToInts(inputArray []string) []int {
	intArrayToReturn := []int{}

	for _, ele := range inputArray { // you can escape index by _ keyword
		if ele == "" {
			// fmt.Println("Found empty string skipping")
			continue
		}

		convertedString, err := strconv.Atoi(ele)

		if err != nil {
			panic(err)
		}

		intArrayToReturn = append(intArrayToReturn, convertedString)
	}
	return intArrayToReturn
}

type PuzzelArray struct {
	data [][]int
}

func make2DArray(boardString string) []int {
	stringsArray := strings.Split(boardString, "\n")

	var c []int
	for _, ele := range strings.Split(strings.Join(stringsArray, " "), " ") {
		// intsArray := strings.Split(v, " ")
		// fmt.Printf("Split array %q\n", intsArray)
		// x := convertArrayToInts(intsArray)
		// fmt.Printf("Int array  %d\n", x)
		if ele == "" {
			// fmt.Println("Found empty string skipping")
			continue
		}

		convertedString, err := strconv.Atoi(ele)

		if err != nil {
			panic(err)
		}
		c = append(c, convertedString)
	}
	// fmt.Printf("Int array  %d\n", c)
	return c
}

func indexOf(arr []int, val int) int {
	for pos, v := range arr {
		if v == val {
			return pos
		}
	}
	return -1
}

func checkWin(b []int) bool {
	win := true
	// let's start with the horizontal rows first
	for i := 0; i < 5; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 5; i < 10; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 10; i < 15; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 15; i < 20; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 20; i < 25; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	// now for the vertical rows
	win = true
	for i := 0; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 1; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}

	win = true
	for i := 2; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 3; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 4; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}

	if win {
		return true
	}

	return false
}

func main() {
	lines := ReadFile("\n\n")

	drawnNumbers := strings.Split(lines[0], ",")
	drawnNumbersInts := convertArrayToInts(drawnNumbers)
	fmt.Printf("drawnNumbersInts%d\n", drawnNumbersInts)

	var boards [][]int

	for i := 1; i < len(lines)-1; i++ {
		board := make2DArray(lines[i])
		// fmt.Printf("Board len  %d\n", len(board))
		boards = append(boards, board)
	}
	// fmt.Printf("Boards winner  %d\n", boards[43])

	// // part 1
	// for _, dn := range drawnNumbersInts {
	// 	for bIndex, b := range boards {
	// 		for index, v := range b {
	// 			if v == dn {
	// 				b[index] = 0
	// 				break
	// 			}
	// 		}

	// 		if checkWin(b) {
	// 			fmt.Println("WIN =>", dn, bIndex, b)
	// 			sum := 0
	// 			for _, num := range b {
	// 				sum += num
	// 			}
	// 			fmt.Println("WIN SUM * NUM =>", sum*dn, sum, dn)
	// 			return
	// 		}
	// 	}
	// }

	boardWin := make([]bool, len(boards))

	// // part 2
	for _, dn := range drawnNumbersInts {
		for b := range boards {
			if boardWin[b] {
				continue
			}

			for index, v := range boards[b] {
				if v == dn {
					boards[b][index] = 0
					break
				}
			}

			if checkWin(boards[b]) {
				// fmt.Println("WIN =>", dn, b)
				sum := 0
				for _, num := range boards[b] {
					sum += num
				}
				fmt.Println("WIN =>", dn, boards[b], sum*dn)
				boardWin[b] = true
			}
		}
	}

}
