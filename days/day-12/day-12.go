package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-12-input.txt"
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

type Cave struct {
	Name  string
	Big   bool
	Other []string
}

func parseInput(input []string) bool {

	caves := map[string]*Cave{}

	for _, edge := range input {

		sp := strings.Split(edge, "-")

		var a, b = sp[0], sp[1]

		if _, ok := caves[a]; !ok {
			caves[a] = &Cave{
				Name:  a,
				Big:   a[0] < 'a',
				Other: []string{b},
			}
		} else {
			caves[a].Other = append(caves[a].Other, b)

		}

		a, b = b, a

		if _, ok := caves[a]; !ok {
			caves[a] = &Cave{
				Name:  a,
				Big:   a[0] < 'a',
				Other: []string{b},
			}
		} else {
			caves[a].Other = append(caves[a].Other, b)

		}
	}

	fmt.Println(caves["end"])
	paths := navigate(caves)
	for _, path := range paths {
		fmt.Println(path)
	}
	fmt.Println(len(paths))
	return true
}

type Path struct {
	Path              []string
	Visited           map[string]bool
	VisitedSmallTwice bool
}

func copyMap(a map[string]bool) map[string]bool {

	out := map[string]bool{}

	for k, v := range a {
		out[k] = v
	}
	return out
}

func navigate(caves map[string]*Cave) [][]string {
	queue := []Path{Path{[]string{"start"}, map[string]bool{}, false}}
	paths := [][]string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		cave := caves[current.Path[len(current.Path)-1]]

		if cave.Name == "end" {
			paths = append(paths, current.Path)
			continue
		}

		newVisited := copyMap(current.Visited)

		if !cave.Big {
			newVisited[cave.Name] = true
		}

		for _, cave := range cave.Other {
			newVisitedSmallTwice := current.VisitedSmallTwice

			if current.Visited[cave] {
				if cave == "start" || cave == "end" || current.VisitedSmallTwice {
					continue
				} else {
					newVisitedSmallTwice = true
				}
			}
			newPath := make([]string, len(current.Path)+1)
			copy(newPath, current.Path)
			newPath = append(newPath, cave)

			queue = append(queue, Path{newPath, newVisited, newVisitedSmallTwice})
		}
	}
	return paths
}

func main() {
	inputs := ReadFile("\n")

	// var data = []string{
	// 	"start-A",
	// 	"start-b",
	// 	"A-c",
	// 	"A-b",
	// 	"b-d",
	// 	"A-end",
	// 	"b-end",
	// }

	parseInput(inputs)
}
