package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-09-input.txt"
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

func isLowestPoint() {

}

func main() {
	inputs := ReadFile("\n")

	var grid [][]int = make([][]int, len(inputs))

	for i, v := range inputs {
		var stringsArray = strings.Split(v, "")
		grid[i] = convertArrayToInts(stringsArray)
	}
	fmt.Println(grid)

	maxX := len(inputs[0])
	maxY := len(inputs)

	risk := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			var current = grid[y][x]

			if y > 0 && grid[y-1][x] <= current {
				// up
				continue
			}

			if y < maxY-1 && grid[y+1][x] <= current {
				// down
				continue
			}

			if x > 0 && grid[y][x-1] <= current {
				// left
				continue
			}
			if x < maxX-1 && grid[y][x+1] <= current {
				// right
				continue
			}
			risk += current + 1
		}
	}

	fmt.Println("This is the risk", risk)

	var grid2 [][]rune = make([][]rune, len(inputs))
	for i, v := range inputs {
		grid2[i] = make([]rune, len(v))
	}

	type Point struct {
		x, y int
	}
	name := 'a'

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid2[y][x] != 0 {
				continue
			}

			if grid[y][x] == 9 {
				grid2[y][x] = '9'
				continue
			}

			queue := []Point{Point{x, y}}

			for len(queue) > 0 {
				point := queue[0]
				queue := queue[1:]

				grid2[point.y][point.x] = name

				if point.y > 0 && grid[point.y-1][point.x] != 9 && grid2[point.y-1][point.x] == 0 {
					queue = append(queue, Point{point.x, point.y - 1})
				}

				if point.y < maxY-1 && grid[point.y+1][point.x] != 9 && grid2[point.y+1][point.x] == 0 {
					queue = append(queue, Point{point.x, point.y + 1})
				}

				if point.x > 0 && grid[point.y][point.x-1] != 9 && grid2[point.y][point.x-1] == 0 {
					queue = append(queue, Point{point.x - 1, point.y})
				}

				if point.x < maxX-1 && grid[point.y][point.x+1] != 9 && grid2[point.y][point.x+1] == 0 {
					queue = append(queue, Point{point.x + 1, point.y})
				}
			}
		}
		name++
	}

	counts := map[rune]int{}

	for y := 0; y < len(grid2); y++ {
		for x := 0; x < len(grid2[y]); x++ {
			if grid2[y][x] == '9' {
				continue
			}

			cur := grid2[y][x]

			counts[cur]++
		}

	}

	fmt.Println("This is the risk2", counts)
}
