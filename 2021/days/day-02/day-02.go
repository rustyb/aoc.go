package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-02-input.txt"
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

func convertArrayToArrays(inputArray []string) [][]string {
	intArrayToReturn := [][]string{}

	for _, ele := range inputArray { // you can escape index by _ keyword
		stringSplit := strings.Split(ele, " ")
		// convertedString, err := strconv.Atoi(ele)

		// if err != nil {
		// 	panic(err)
		// }

		intArrayToReturn = append(intArrayToReturn, stringSplit)
	}
	return intArrayToReturn
}

func makeInt(value string) int {
	convertedValue, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return convertedValue
}

func main() {
	stringArray := ReadFile("\n")
	splitArray := convertArrayToArrays(stringArray)

	horizontalTotal := 0
	verticalTotal := 0
	aim := 0

	for i := 0; i < len(splitArray); i++ {
		instruction := splitArray[i][0]
		value := makeInt(splitArray[i][1])

		if instruction == "forward" {
			verticalTotal = verticalTotal + (value * aim)
			horizontalTotal += value
		}

		if instruction == "down" {
			// verticalTotal += value
			aim += value
		}

		if instruction == "up" {
			// verticalTotal -= value
			aim -= value
		}
	}
	println("AIM", aim)
	fmt.Printf("Part 1: verticaltotal %d, horizontalTotal %d, multi %d \n", verticalTotal, horizontalTotal, verticalTotal*horizontalTotal)
}
