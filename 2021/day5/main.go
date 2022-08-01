package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parse(str string) (int, int, int, int) {
	coordinates := strings.Split(str, " -> ")
	start := strings.Split(coordinates[0], ",")
	end := strings.Split(coordinates[1], ",")
	x1, _ := strconv.Atoi(start[0])
	y1, _ := strconv.Atoi(start[1])
	x2, _ := strconv.Atoi(end[0])
	y2, _ := strconv.Atoi(end[1])
	return x1, x2, y1, y2

}
func part1(x1 int, x2 int, y1 int, y2 int, board map[int]map[int]int) map[int]map[int]int {
	if x1 == x2 || y1 == y2 {
		if math.Abs(float64(x2-x1)) > 0 {
			var i int = x1
			var end = x2
			if x1 > x2 {
				i = x2
				end = x1

			}

			for i <= end {
				if _, found := board[i]; !found {
					board[i] = map[int]int{}
				}
				board[i][y1] += 1
				i++
			}
		}

		if math.Abs(float64(y2-y1)) > 0 {
			i := y1
			end := y2
			if y1 > y2 {
				i = y2
				end = y1
			}
			for i <= end {
				if _, found := board[x1]; !found {
					board[x1] = map[int]int{}
				}
				board[x1][i] += 1
				i++
			}
		}
	}

	return board
}

func part2(x1 int, y1 int, x2 int, y2 int, board map[int]map[int]int) map[int]map[int]int {

	if x1 != x2 && y1 != y2 {
		if _, found := board[x1]; !found {
			board[x1] = map[int]int{}
		}
		board[x1][y1] += 1

		for x1 != x2 && y1 != y2 {
			if x2 > x1 {
				x1++
			} else {
				x1--
			}
			if y2 > y1 {
				y1++
			} else {
				y1--
			}
			if _, found := board[x1]; !found {
				board[x1] = map[int]int{}
			}
			board[x1][y1] += 1
		}
	}

	return board
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	board := make(map[int]map[int]int)
	for scanner.Scan() {

		currentValue := scanner.Text()
		x1, x2, y1, y2 := parse(currentValue)
		// part 1
		board = part1(x1, x2, y1, y2, board)
		// part 2
		board = part2(x1, y1, x2, y2, board)
	}
	if err != nil {
		fmt.Println("error", err)
	}
	max := 0
	for _, v := range board {
		for _, vv := range v {
			if vv > 1 {
				max++
			}
		}
	}
	fmt.Println("max", max)
	// part1(card, numbers)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
