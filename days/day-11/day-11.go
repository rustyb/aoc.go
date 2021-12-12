package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-11-input.txt"
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

// func step()

var data = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526",
}

func step(grid [][]int) int {

	flashed := [10][10]bool{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x]++
		}
	}

	updatedSomething := true

	for updatedSomething {
		updatedSomething = false

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {

				if flashed[y][x] {
					continue
				}

				if grid[y][x] > 9 {
					updatedSomething = true
					flashed[y][x] = true

					if y-1 >= 0 {
						if x-1 >= 0 {
							grid[y-1][x-1]++
						}
						grid[y-1][x]++
						if x+1 < len(grid[y]) {
							grid[y-1][x+1]++
						}
					}

					if x-1 >= 0 {
						grid[y][x-1]++
					}
					if x+1 < len(grid[y]) {
						grid[y][x+1]++
					}

					if y+1 < len(grid) {
						if x-1 >= 0 {
							grid[y+1][x-1]++
						}
						grid[y+1][x]++
						if x+1 < len(grid[y]) {
							grid[y+1][x+1]++
						}
					}

				}
			}
		}
	}

	count := 0
	fmt.Println(grid)

	for y := 0; y < len(flashed); y++ {
		for x := 0; x < len(flashed[y]); x++ {
			if flashed[y][x] {
				grid[y][x] = 0
				count++
			}
		}
	}

	return count

}

func main() {
	inputs := ReadFile("\n")

	var grid [][]int = make([][]int, len(inputs))
	for i, octo := range inputs {
		var stringsArray = strings.Split(octo, "")
		grid[i] = convertArrayToInts(stringsArray)
	}

	// Part 1
	fmt.Println(grid)
	totalFlashes := 0
	for r := 0; r < 100; r++ {
		fmt.Println("ROUND ", r)
		var flashes = step(grid)
		fmt.Println(flashes)
		totalFlashes += flashes
	}

	fmt.Println("TOTAL FLASHES", totalFlashes)
	// fmt.Println(step(grid))

	// Part 2
	inputs1 := ReadFile("\n")
	var gridP2 [][]int = make([][]int, len(inputs1))
	for i, octo := range inputs1 {
		var stringsArray = strings.Split(octo, "")
		gridP2[i] = convertArrayToInts(stringsArray)
	}

	fmt.Println(gridP2)
	totalFlashes = 0
	for r := 0; r < 1000; r++ {
		fmt.Println("ROUND ", r)
		var flashes = step(gridP2)
		if flashes == 100 {
			fmt.Println("simultaneously =>", r+1)
			break
		}
		fmt.Println(flashes)
		totalFlashes += flashes
	}

}
