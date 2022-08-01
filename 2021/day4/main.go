package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
try to put in a same array
the vertical line and the horizontal line
to be able to look up easier
**/
type cell struct {
	value     int
	bingo     bool
	direction string
	index     int
}

func addVertical(line []cell) []cell {
	customIndex := 0
	for i := 0; i < len(line)/5; i++ {
		for _, v := range line {
			if i == v.index {
				line = append(line, cell{
					value:     v.value,
					bingo:     false,
					direction: "vertical",
					index:     customIndex,
				})
				if customIndex == 4 {
					customIndex = 0
				}

				customIndex++

			}

		}
	}
	return line
}
func parse() {
}

func part1(card [][]cell, numbers []string) {
	numberJustCalled := 0
	var ranking [][]cell
	var numberCalledRanking []int
out:
	for _, n := range numbers {
		fmt.Println("card len", len(card))
		for m := 0; m < len(card); m++ {
			for i := 0; i < len(card[m]); i++ {
				nIint, _ := strconv.Atoi(n)
				if card[m][i].direction == "horizo" && nIint == card[m][i].value {
					card[m][i].bingo = true
				}
				if card[m][i].direction == "vertical" && nIint == card[m][i].value {
					card[m][i].bingo = true
				}
			}
			winH := 0
			winV := 0
			for j := 0; j < len(card[m]); j++ {
				if card[m][j].direction == "horizo" && card[m][j].bingo == true {
					winH++
				}
				if card[m][j].direction == "vertical" && card[m][j].bingo == true {
					winV++
				}
				if winH == 5 || winV == 5 {
					ranking = append(ranking, card[m])
					numberCalled, _ := strconv.Atoi(n)
					numberJustCalled = numberCalled
					numberCalledRanking = append(numberCalledRanking, numberJustCalled)
					/**
					* added to retrieve the worst board
					**/
					if len(card) == 1 {
						break out
					} else if m < len(card) {
						card = append(card[:m], card[m+1:]...)
						m = 0
						break
					}
				}
				if card[m][j].index == 4 && card[m][j].direction == "horizo" {
					winH = 0
				}
				if card[m][j].index == 4 && card[m][j].direction == "vertical" {
					winV = 0
				}
			}

		}

	}
	sumNotCalled := 0
	for _, v := range ranking[len(ranking)-1] {

		if v.bingo == false && v.direction == "horizo" {
			sumNotCalled += v.value

		}
	}
	fmt.Println("numberJustCAlled", numberCalledRanking[len(numberCalledRanking)-1])
	fmt.Println("sumNotCalled", sumNotCalled)
	fmt.Println("score", numberJustCalled*sumNotCalled)
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var card [][]cell
	var numbers []string

	var line []cell
	for scanner.Scan() {

		currentValue := scanner.Text()
		if len(numbers) == 0 {
			numbers = strings.Split(currentValue, ",")
			continue
		}
		if len(currentValue) == 0 {
			if len(line) > 0 {
				line = addVertical(line)
				card = append(card, line)
				line = make([]cell, 0)
			}
			continue
		}

		values := strings.Fields(currentValue)
		for k, v := range values {
			v := strings.Trim(v, " ")
			if v != "" {
				vInt, _ := strconv.Atoi(v)
				line = append(line, cell{
					value:     vInt,
					bingo:     false,
					direction: "horizo",
					index:     k,
				})

			}
		}

	}

	line = addVertical(line)

	card = append(card, line)
	if err != nil {
		fmt.Println("error", err)
	}
	part1(card, numbers)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
