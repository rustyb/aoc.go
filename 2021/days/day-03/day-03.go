package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-03-input.txt"
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

func mostAndLeastCommonValue(lines []string, i int) (uint8, uint8) {
	count0, count1 := 0, 0
	for _, line := range lines {
		if line[i] == '0' {
			count0++
		} else {
			count1++
		}
	}
	if count0 > count1 {
		return '0', '1'
	}
	return '1', '0'
}

func convertBinaryToInt(binaryString string) int64 {
	i, err := strconv.ParseInt("1101", 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func fromBinary(s string) (r int) {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func filter(lines []string, i int, mostCommon bool) string {
	if len(lines) == 1 {
		return lines[0]
	}
	most, least := mostAndLeastCommonValue(lines, i)
	comparator := least
	if mostCommon {
		comparator = most
	}
	filtered := make([]string, 0)
	for _, l := range lines {
		if l[i] == comparator {
			filtered = append(filtered, l)
		}
	}
	return filter(filtered, i+1, mostCommon)
}

func main() {
	lines := ReadFile("\n")

	gamma, epsilon := "", ""
	for i := 0; i < len(lines[0]); i++ {
		most, least := mostAndLeastCommonValue(lines, i)
		gamma += string(most)
		epsilon += string(least)
	}
	fmt.Println(fromBinary(gamma) * fromBinary(epsilon))

	oxygen := filter(lines, 0, true)
	scrubber := filter(lines, 0, false)
	fmt.Println(fromBinary(oxygen) * fromBinary(scrubber))
}
