package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-01-input.txt"
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	fmt.Println(fileContent)
	stringSplit := strings.Split(fileContent, delimeter)

	// To deal with the \n on the end of the file read in
	// fmt.Printf("%q\n", stringSplit[:len(stringSplit)-1])
	return stringSplit[:len(stringSplit)-1]
}

// convert the string array to ints to be able to
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

func lessThanPreviousValue(value int, previousValue int) bool {
	return value > previousValue
}

func main() {
	stringArray := ReadFile("\n")
	intArray := convertArrayToInts(stringArray)
	countDecreasing := 0

	for i := 0; i < len(intArray); i++ {
		if i > 0 {
			isDecreasing := lessThanPreviousValue(intArray[i], intArray[i-1])
			if isDecreasing == true {
				countDecreasing += 1
			}
		}
	}

	part2IsDecreasing := 0
	previousThreeValueSum := 0
	chunkSize := 3

	for i := 0; i < len(intArray); i++ {
		end := i + chunkSize
		if end <= len(intArray) {
			sum := intArray[i] + intArray[i+1] + intArray[i+2]
			println("END => ", end, "SUM => ", sum, "Prev sum => ", previousThreeValueSum, "A A A", intArray[i], intArray[i+1], intArray[i+2])
			if previousThreeValueSum > 0 {
				isDecreasing := lessThanPreviousValue(sum, previousThreeValueSum)
				if isDecreasing == true {
					part2IsDecreasing += 1
				}
			}
			previousThreeValueSum = sum
		}
	}

	fmt.Printf("Part 1: %d results are larger than previous measurement\n", countDecreasing)
	fmt.Printf("Part 2: %d results are larger than previous measurement\n", part2IsDecreasing)
}
