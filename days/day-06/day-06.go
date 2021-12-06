package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-06-input.txt"
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

func advanceTime(fish []int) []int {
	var newFish []int
	// origFishPlus := make([]int, len(fish))

	for i := range fish {
		if fish[i] == 0 {
			newFish = append(newFish, 8)
			// origFishPlus = append(origFishPlus, 7)
			fish[i] = 6
		} else {
			fish[i]--
		}
	}
	// for i := 0; i < len(fish); i++ {
	// 	// newEnergy := fish[i] - 1

	// 	if fish[i] == 0 {
	// 		newFish = append(newFish, 8)
	// 		// origFishPlus = append(origFishPlus, 7)
	// 		origFishPlus[i] = 6
	// 	} else {
	// 		// origFishPlus = append(origFishPlus, newEnergy)
	// 		origFishPlus[i] = fish[i] - 1
	// 	}
	// }

	fish = append(fish, newFish...)
	// fmt.Printf("combined fish %d\n", origFishPlus)
	return fish
}

func advanceTimePart2(fish []int) []int {
	var next = make([]int, 9)
	for i := 1; i < 9; i++ {
		next[i-1] = fish[i]
	}

	next[6] += fish[0]
	next[8] += fish[0]
	return next
}

func part2(input []int) []int {
	fishes := make([]int, len(input))
	copy(fishes, input)

	// let's only store the fish by there age eg 0 - 8
	var fish1 = make([]int, 9)
	for i := range fishes {
		fish1[fishes[i]]++
	}

	fmt.Println("Fish1 before", fish1)
	for i := 0; i < 256; i++ {
		fish1 = advanceTimePart2(fish1)
	}
	fmt.Println("Fish1 after", fish1)

	return fish1
}

func main() {
	input := ReadFile("\n")
	input = strings.Split(input[0], ",")
	fmt.Println(convertArrayToInts(input))
	asNumbers := convertArrayToInts(input)

	nextDay := asNumbers

	for i := 0; i < 80; i++ {
		nextDay = advanceTime(nextDay)
	}
	fmt.Println("Length after 80 => ", len(nextDay))

	input2 := ReadFile("\n")
	input2 = strings.Split(input2[0], ",")
	asNumbers2 := convertArrayToInts(input2)
	day2Count := part2(asNumbers2)

	sum := 0
	for _, v := range day2Count {
		sum += v
	}

	fmt.Println("Length after 256", sum)

}
