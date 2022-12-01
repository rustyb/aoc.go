package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-05-input.txt"
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

type Point struct {
	x int
	y int
}

func pointFromString(input string) (*Point, error) {
	sp := strings.Split(input, ",")
	x, err := strconv.Atoi(sp[0])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(sp[1])
	if err != nil {
		return nil, err
	}
	return &Point{x: x, y: y}, nil
}

type Line [2]*Point

func parse(input []string) ([]Line, error) {
	var o []Line
	for _, v := range input {
		points := strings.Split(v, " -> ")
		p1, err := pointFromString(points[0])

		if err != nil {
			return nil, err
		}

		p2, err := pointFromString(points[1])

		if err != nil {
			return nil, err
		}
		o = append(o, Line{p1, p2})
	}
	return o, nil
}

func removeDiagonalLines(lines []Line) []Line {
	var n int
	for _, line := range lines {
		if line[0].x == line[1].x || line[0].y == line[1].y {
			lines[n] = line
			n += 1
		}
	}
	return lines[:n]
}

func iteratePoints(line Line) chan *Point {

	p1 := line[0]
	p2 := line[1]

	delta_x := p2.x - p1.x
	delta_y := p2.y - p1.y

	var xStep, yStep int

	if delta_x > 0 {
		xStep = 1
	} else if delta_x < 0 {
		xStep = -1
	}

	if delta_y > 0 {
		yStep = 1
	} else if delta_y < 0 {
		yStep = -1
	}

	c := make(chan *Point)

	go func() {
		c <- p1
		lastPoint := *p1
		for lastPoint != *p2 {
			np := &Point{
				x: lastPoint.x + xStep,
				y: lastPoint.y + yStep,
			}
			c <- np
			lastPoint = *np
		}
		close(c)
	}()

	return c
}

func countOverlappingPoints(lines []Line) int {
	areas := make(map[Point]int)
	for _, line := range lines {
		for point := range iteratePoints(line) {
			t := *point
			areas[t] = areas[t] + 1
		}
	}

	var numberOverlaps int
	for _, n := range areas {
		if n > 1 {
			numberOverlaps += 1
		}
	}
	return numberOverlaps
}

func main() {
	input := ReadFile("\n")
	lines, err := parse(input)
	if err != nil {
		fmt.Printf("ERROR %q\n", err)
		return
	}

	linesr := removeDiagonalLines(lines)

	fmt.Println("overlapping point count!", countOverlappingPoints(linesr))

	// fmt.Printf("Split array %q\n", lines)

	// Part 2
	input2 := ReadFile("\n")
	lines2, err2 := parse(input2)
	if err2 != nil {
		fmt.Printf("ERROR %q\n", err2)
		return
	}
	fmt.Println("overlapping point count2!", countOverlappingPoints(lines2))
}
