package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type minCoord struct {
	x     int
	y     int
	value string
}

type adjacentNumber struct {
	x        int
	y        int
	value    string
	neighbor []*adjacentNumber
}

func topLeftCornerNeighbor(x int, y int, heightmap [][]string) (adjacentNumber, adjacentNumber) {
	n := adjacentNumber{
		x:        x,
		y:        y + 1,
		value:    heightmap[y+1][x],
		neighbor: []*adjacentNumber{},
	}
	n1 := adjacentNumber{
		x:     x + 1,
		y:     y,
		value: heightmap[y][x+1],

		neighbor: []*adjacentNumber{},
	}
	return n, n1
}
func topRightCornerNeighbor(x int, y int, heightmap [][]string) (adjacentNumber, adjacentNumber) {
	n := adjacentNumber{
		x:        x,
		y:        y + 1,
		value:    heightmap[y+1][x],
		neighbor: []*adjacentNumber{},
	}
	n1 := adjacentNumber{
		x:     x - 1,
		y:     y,
		value: heightmap[y][x-1],

		neighbor: []*adjacentNumber{},
	}

	return n, n1
}
func bottomRightCornerNeighbor(x int, y int, heightmap [][]string) (adjacentNumber, adjacentNumber) {
	n := adjacentNumber{
		x:        x,
		y:        y - 1,
		value:    heightmap[y-1][x],
		neighbor: []*adjacentNumber{},
	}
	n1 := adjacentNumber{
		x:     x - 1,
		y:     y,
		value: heightmap[y][x-1],

		neighbor: []*adjacentNumber{},
	}
	return n, n1
}
func bottomLeftCornerNeighbor(x int, y int, heightmap [][]string) (adjacentNumber, adjacentNumber) {
	n := adjacentNumber{
		x:        x,
		y:        y - 1,
		value:    heightmap[y-1][x],
		neighbor: []*adjacentNumber{},
	}
	n1 := adjacentNumber{
		x:     x + 1,
		y:     y,
		value: heightmap[y][x+1],

		neighbor: []*adjacentNumber{},
	}
	return n, n1
}
func leftBorderNeighbor(x int, y int, heightmap [][]string) (adjacentNumber, adjacentNumber, adjacentNumber) {
	n := adjacentNumber{
		x:        x,
		y:        y - 1,
		value:    heightmap[y-1][x],
		neighbor: []*adjacentNumber{},
	}
	n1 := adjacentNumber{
		x:        x,
		y:        y + 1,
		value:    heightmap[y+1][x],
		neighbor: []*adjacentNumber{},
	}
	n2 := adjacentNumber{
		x:     x + 1,
		y:     y,
		value: heightmap[y][x+1],

		neighbor: []*adjacentNumber{},
	}
	return n, n1, n2
}
func rightBorderNeighbord(x int, y int, heightmap [][]string) (adjacentNumber, adjacentNumber, adjacentNumber) {
	n := adjacentNumber{
		x:        x,
		y:        y - 1,
		value:    heightmap[y-1][x],
		neighbor: []*adjacentNumber{},
	}
	n1 := adjacentNumber{
		x:        x,
		y:        y + 1,
		value:    heightmap[y+1][x],
		neighbor: []*adjacentNumber{},
	}
	n2 := adjacentNumber{
		x:     x - 1,
		y:     y,
		value: heightmap[y][x-1],

		neighbor: []*adjacentNumber{},
	}
	return n, n1, n2

}
func bottomBorderNeighbord(x int, y int, heightmap [][]string) (adjacentNumber, adjacentNumber, adjacentNumber) {
	n := adjacentNumber{
		x:        x,
		y:        y - 1,
		value:    heightmap[y-1][x],
		neighbor: []*adjacentNumber{},
	}
	n1 := adjacentNumber{
		x:        x + 1,
		y:        y,
		value:    heightmap[y][x+1],
		neighbor: []*adjacentNumber{},
	}
	n2 := adjacentNumber{
		x:        x - 1,
		y:        y,
		value:    heightmap[y][x-1],
		neighbor: []*adjacentNumber{},
	}
	return n, n1, n2
}
func topBorderNeighbord(x int, y int, heightmap [][]string) (adjacentNumber, adjacentNumber, adjacentNumber) {
	n := adjacentNumber{
		x:        x,
		y:        y + 1,
		value:    heightmap[y+1][x],
		neighbor: []*adjacentNumber{},
	}
	n1 := adjacentNumber{
		x:        x + 1,
		y:        y,
		value:    heightmap[y][x+1],
		neighbor: []*adjacentNumber{},
	}
	n2 := adjacentNumber{
		x:     x - 1,
		y:     y,
		value: heightmap[y][x-1],

		neighbor: []*adjacentNumber{},
	}
	return n, n1, n2
}
func classicNeighbord(x int, y int, heightmap [][]string) (adjacentNumber, adjacentNumber, adjacentNumber, adjacentNumber) {
	n := adjacentNumber{
		x:        x,
		y:        y + 1,
		value:    heightmap[y+1][x],
		neighbor: []*adjacentNumber{},
	}
	n1 := adjacentNumber{
		x:        x,
		y:        y - 1,
		value:    heightmap[y-1][x],
		neighbor: []*adjacentNumber{},
	}
	n2 := adjacentNumber{
		x:        x + 1,
		y:        y,
		value:    heightmap[y][x+1],
		neighbor: []*adjacentNumber{},
	}
	n3 := adjacentNumber{
		x:     x - 1,
		y:     y,
		value: heightmap[y][x-1],

		neighbor: []*adjacentNumber{},
	}
	return n, n1, n2, n3
}
func part1(heightmap [][]string) []minCoord {
	var min []int
	var minWithLocation []minCoord
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			// top left corner
			if j == 0 && i == 0 {
				if heightmap[i][j] < heightmap[i][j+1] && heightmap[i][j] < heightmap[i+1][j] {
					intDigit, _ := strconv.Atoi(heightmap[i][j])
					min = append(min, intDigit)

					minWithLocation = append(minWithLocation, minCoord{x: j, y: i, value: heightmap[i][j]})
				}
			}
			// top right corner
			if j == len(heightmap[i])-1 && i == 0 {
				if heightmap[i][j] < heightmap[i+1][j] && heightmap[i][j] < heightmap[i][j-1] {

					intDigit, _ := strconv.Atoi(heightmap[i][j])
					min = append(min, intDigit)

					minWithLocation = append(minWithLocation, minCoord{x: j, y: i, value: heightmap[i][j]})
				}
			}
			if j == len(heightmap[i])-1 && i > 0 && i < len(heightmap)-1 {
				if heightmap[i][j] < heightmap[i+1][j] && heightmap[i][j] < heightmap[i][j-1] && heightmap[i][j] < heightmap[i-1][j] {

					intDigit, _ := strconv.Atoi(heightmap[i][j])
					min = append(min, intDigit)

					minWithLocation = append(minWithLocation, minCoord{x: j, y: i, value: heightmap[i][j]})
				}
			}
			// bottom left corner
			if i == len(heightmap)-1 && j == 0 {
				if heightmap[i][j] < heightmap[i][j+1] && heightmap[i][j] < heightmap[i-1][j] {

					intDigit, _ := strconv.Atoi(heightmap[i][j])
					min = append(min, intDigit)

					minWithLocation = append(minWithLocation, minCoord{x: j, y: i, value: heightmap[i][j]})
				}

			}
			if i == 0 && j > 0 && j < len(heightmap[i])-1 {

				if heightmap[i][j] < heightmap[i+1][j] && heightmap[i][j] < heightmap[i][j+1] && heightmap[i][j] < heightmap[i][j-1] {
					intDigit, _ := strconv.Atoi(heightmap[i][j])
					min = append(min, intDigit)

					minWithLocation = append(minWithLocation, minCoord{x: j, y: i, value: heightmap[i][j]})
				}

			}
			if j == 0 && i < len(heightmap)-1 && i > 0 {

				if heightmap[i][j] < heightmap[i+1][j] && heightmap[i][j] < heightmap[i][j+1] && heightmap[i][j] < heightmap[i-1][j] {

					intDigit, _ := strconv.Atoi(heightmap[i][j])
					min = append(min, intDigit)

					minWithLocation = append(minWithLocation, minCoord{x: j, y: i, value: heightmap[i][j]})
				}

			}
			if i == len(heightmap)-1 && j > 0 && j < len(heightmap[i])-1 {

				if heightmap[i][j] < heightmap[i-1][j] && heightmap[i][j] < heightmap[i][j+1] && heightmap[i][j] < heightmap[i][j-1] {
					intDigit, _ := strconv.Atoi(heightmap[i][j])
					min = append(min, intDigit)

					minWithLocation = append(minWithLocation, minCoord{x: j, y: i, value: heightmap[i][j]})
				}

			}
			// bottom right corner
			if i == len(heightmap)-1 && j == len(heightmap[i])-1 {
				if heightmap[i][j] < heightmap[i][j-1] && heightmap[i][j] < heightmap[i-1][j] {

					intDigit, _ := strconv.Atoi(heightmap[i][j])
					min = append(min, intDigit)

					minWithLocation = append(minWithLocation, minCoord{x: j, y: i, value: heightmap[i][j]})
				}

			}
			// else
			if i > 0 && j > 0 && i < len(heightmap)-1 && j < len(heightmap[i])-1 {
				if heightmap[i][j] < heightmap[i][j-1] &&
					heightmap[i][j] < heightmap[i][j+1] &&
					heightmap[i][j] < heightmap[i-1][j] &&
					heightmap[i][j] < heightmap[i+1][j] {
					intDigit, _ := strconv.Atoi(heightmap[i][j])
					min = append(min, intDigit)
					minWithLocation = append(minWithLocation, minCoord{x: j, y: i, value: heightmap[i][j]})
				}

			}
		}
	}
	fmt.Println("min", min)
	sum := 0
	for _, v := range min {
		sum += v + 1
	}
	fmt.Println("sum", sum)
	return minWithLocation
}
func buildTree(heightmap [][]string) ([]adjacentNumber, map[string]adjacentNumber) {
	/*
	*
	* build a tree with the curent number and its neighbors
	* the key is a string defined with i + j
	* becareful if i or j is a single digit, add a 0 to avoid confusion
	*
	* could be improved by building the tree with recursion
	* and looking for neighbors of neighbors
	 */
	var result []adjacentNumber
	result1 := make(map[string]adjacentNumber)
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			main := adjacentNumber{
				x:        j,
				y:        i,
				value:    heightmap[i][j],
				neighbor: []*adjacentNumber{},
			}
			iStr := strconv.Itoa(i)
			jStr := strconv.Itoa(j)
			if len(iStr) == 1 {
				iStr = "0" + iStr
			}
			if len(jStr) == 1 {
				jStr = "0" + jStr
			}
			result1[jStr+iStr] = adjacentNumber{
				x:        j,
				y:        i,
				value:    heightmap[i][j],
				neighbor: []*adjacentNumber{},
			}

			if i == 0 && j == 0 {
				n, n1 := topLeftCornerNeighbor(j, i, heightmap)
				main.neighbor = append(main.neighbor, &n, &n1)
				if entry, ok := result1[jStr+iStr]; ok {
					entry.neighbor = append(entry.neighbor, &n, &n1)
					result1[jStr+iStr] = entry
				}

			}

			if j == len(heightmap[i])-1 && i == 0 {
				n, n1 := topRightCornerNeighbor(j, i, heightmap)
				main.neighbor = append(main.neighbor, &n, &n1)
				if entry, ok := result1[jStr+iStr]; ok {
					entry.neighbor = append(entry.neighbor, &n, &n1)
					result1[jStr+iStr] = entry
				}
			}

			if j == len(heightmap[i])-1 && i > 0 && i < len(heightmap)-1 {
				n, n1, n2 := rightBorderNeighbord(j, i, heightmap)
				main.neighbor = append(main.neighbor, &n, &n1, &n2)
				if entry, ok := result1[jStr+iStr]; ok {
					entry.neighbor = append(entry.neighbor, &n, &n1, &n2)
					result1[jStr+iStr] = entry
				}
			}

			if i == len(heightmap)-1 && j == 0 {
				n, n1 := bottomLeftCornerNeighbor(j, i, heightmap)
				main.neighbor = append(main.neighbor, &n, &n1)
				if entry, ok := result1[jStr+iStr]; ok {
					entry.neighbor = append(entry.neighbor, &n, &n1)
					result1[jStr+iStr] = entry
				}
			}

			if i == 0 && j > 0 && j < len(heightmap[i])-1 {
				fmt.Println("i=", i, "j=", j)
				n, n1, n2 := topBorderNeighbord(j, i, heightmap)
				main.neighbor = append(main.neighbor, &n, &n1, &n2)
				if entry, ok := result1[jStr+iStr]; ok {
					entry.neighbor = append(entry.neighbor, &n, &n1, &n2)
					result1[jStr+iStr] = entry
				}
			}

			if j == 0 && i < len(heightmap)-1 && i > 0 {
				n, n1, n2 := leftBorderNeighbor(j, i, heightmap)
				main.neighbor = append(main.neighbor, &n, &n1, &n2)
				if entry, ok := result1[jStr+iStr]; ok {
					entry.neighbor = append(entry.neighbor, &n, &n1, &n2)
					result1[jStr+iStr] = entry
				}
			}

			if i == len(heightmap)-1 && j > 0 && j < len(heightmap[i])-1 {
				n, n1, n2 := bottomBorderNeighbord(j, i, heightmap)
				main.neighbor = append(main.neighbor, &n, &n1, &n2)
				if entry, ok := result1[jStr+iStr]; ok {
					entry.neighbor = append(entry.neighbor, &n, &n1, &n2)
					result1[jStr+iStr] = entry
				}
			}

			if i == len(heightmap)-1 && j == len(heightmap[i])-1 {
				n, n1 := bottomRightCornerNeighbor(j, i, heightmap)
				main.neighbor = append(main.neighbor, &n, &n1)
				if entry, ok := result1[jStr+iStr]; ok {
					entry.neighbor = append(entry.neighbor, &n, &n1)
					result1[jStr+iStr] = entry
				}
			}

			if i > 0 && j > 0 && i < len(heightmap)-1 && j < len(heightmap[i])-1 {
				n, n1, n2, n3 := classicNeighbord(j, i, heightmap)
				main.neighbor = append(main.neighbor, &n, &n1, &n2, &n3)
				if entry, ok := result1[jStr+iStr]; ok {
					entry.neighbor = append(entry.neighbor, &n, &n1, &n2, &n3)
					result1[jStr+iStr] = entry
				}
			}
			result = append(result, main)
		}
	}
	return result, result1
}
func toTheTail(n adjacentNumber) {
	if len(n.neighbor) == 0 {
		fmt.Println("value", n.value)
		fmt.Println("end")
	} else {
		for _, v := range n.neighbor {
			toTheTail(*v)
		}
	}
}
func toTheNine(n adjacentNumber, result1 map[string]adjacentNumber, visited map[string]int) {
	/**
	* 9 is the exit number
	 */
	if n.value != "9" {
		for _, v := range n.neighbor {
			xStr := strconv.Itoa(v.x)
			yStr := strconv.Itoa(v.y)
			if len(yStr) == 1 {
				yStr = "0" + yStr
			}
			if len(xStr) == 1 {
				xStr = "0" + xStr
			}

			if _, found := visited[xStr+yStr]; !found && result1[xStr+yStr].value != "9" {
				visited[xStr+yStr] = 1
				toTheNine(result1[xStr+yStr], result1, visited)

			}
		}
	}
}
func part2(heightmap [][]string, coord []minCoord) {
	/**
	* build the linked list tree
	* by adding the neighbors as a children of the main number
	*
	 */
	_, result1 := buildTree(heightmap)
	fmt.Println("test", result1["0000"].neighbor)
	fmt.Println("test", result1["0000"].neighbor[0].neighbor[0])
	var threeLargest []int
	for _, v := range coord {

		xStr := strconv.Itoa(v.x)
		yStr := strconv.Itoa(v.y)
		if len(xStr) == 1 {
			xStr = "0" + xStr
		}
		if len(yStr) == 1 {
			yStr = "0" + yStr
		}

		visited := map[string]int{xStr + yStr: 1}
		toTheNine(result1[xStr+yStr], result1, visited)
		threeLargest = append(threeLargest, len(visited))
	}
	sort.Ints(threeLargest)
	fmt.Println("final fucking result", threeLargest[len(threeLargest)-1]*threeLargest[len(threeLargest)-2]*threeLargest[len(threeLargest)-3])
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	var heightmap [][]string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		currentValue := scanner.Text()
		numbers := strings.Split(currentValue, "")
		heightmap = append(heightmap, numbers)

	}
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("max", heightmap)
	minWithLocation := part1(heightmap)
	part2(heightmap, minWithLocation)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
