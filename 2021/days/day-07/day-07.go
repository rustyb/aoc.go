package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-07-input.txt"
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

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func calculateFuelToPosition(crabs []int, horizontalPosition int) int {
	var fuelRequired = 0
	for _, c := range crabs {
		fuelRequired += Abs(c - horizontalPosition)
	}

	return fuelRequired
}

func cost(i int) int {
	return (i * (i + 1)) / 2
}

func calculateFuelToPositionPart2(crabs []int, horizontalPosition int) int {
	var fuelRequired = 0

	for _, c := range crabs {
		totalDistance := Abs(c - horizontalPosition)
		fuelRequired += cost(totalDistance)
	}

	return fuelRequired
}

func main() {
	inputs := ReadFile("\n")
	numbers := strings.Split(inputs[0], ",")
	crabs := convertArrayToInts(numbers)

	fuelGuesses := []int{}
	for i := 0; i < 2000; i++ {
		fuelGuesses = append(fuelGuesses, calculateFuelToPosition(crabs, i))
	}

	min, max := MinMax(fuelGuesses)
	fmt.Printf("Part 1MIN, MAX => %d %d\n", min, max)

	fuelGuessesP2 := []int{}
	for i := 0; i < 2000; i++ {
		fuelGuessesP2 = append(fuelGuessesP2, calculateFuelToPositionPart2(crabs, i))
	}

	minP2, maxP2 := MinMax(fuelGuessesP2)
	fmt.Printf("Part 1MIN, MAX => %d %d\n", minP2, maxP2)
	// fmt.Printf("%d\n", fuelGuessesP2)
}
