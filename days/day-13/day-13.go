package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-13-input.txt"
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

func convertArrayToInts(inputArray []string) []int {
	intArrayToReturn := []int{}

	for _, ele := range inputArray { // you can escape index by _ keyword
		convertedString, err := strconv.Atoi(ele)

		if err != nil {
			panic(err)
		}

		intArrayToReturn = append(intArrayToReturn, convertedString)
	}
	return intArrayToReturn
}

type stringSlice []string

func (slice stringSlice) pos(value string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func main() {
	inputs := ReadFile("\n")
	// endOfCoords := inputs.pos("")
	// fmt.Println("end of coordinates is ", endOfCoords)

	type xy struct {
		x, y int
	}

	var points []xy
	var folds []xy

	parseFold := false
	for _, row := range inputs {
		fmt.Println("LINE", row, parseFold)
		if row == "" {
			parseFold = true
			continue
		}

		if !parseFold {
			var xy1 xy
			fmt.Sscanf(row, "%d,%d", &xy1.x, &xy1.y)
			points = append(points, xy1)
		} else {
			var dir rune
			var value int
			fmt.Sscanf(row, "fold along %c=%d", &dir, &value)
			if dir == 'x' {
				folds = append(folds, xy{x: value})
			} else {
				folds = append(folds, xy{y: value})
			}
		}
	}
	fmt.Println("Points", points)
	fmt.Println("Folds", folds)

	// maxX := folds[0].x * 2
	// maxY := folds[1].y * 2

	var grid [][]int = make([][]int, 6)
	for i := range grid {
		grid[i] = make([]int, 40)
	}

	for _, p := range points {
		var x, y int = p.x, p.y

		for _, fold := range folds {
			if fold.x > 0 {
				if x > fold.x {
					x = 2*fold.x - x

				}
			} else {
				if y > fold.y {
					y = 2*fold.y - y
				}
			}
		}

		grid[y][x] = 1
	}

	for _, row := range grid {
		for _, v := range row {
			if v == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
