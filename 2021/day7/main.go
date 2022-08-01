package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(fishes *[]int) {
}

func part2(fishes *map[int]int) {
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var positions []int
	for scanner.Scan() {

		currentValue := scanner.Text()
		positionStr := strings.Split(currentValue, ",")
		for _, v := range positionStr {
			vInt, _ := strconv.Atoi(v)
			positions = append(positions, vInt)
		}

	}
	min := 999999999999999999
	for i := 0; i < len(positions); i++ {
		var currentMin int
		for _, v := range positions {
			result := int(math.Abs(float64(v - positions[i])))
			currentMin += result
		}
		if currentMin < min {
			min = currentMin
		}
	}

	fmt.Println("min", min)
	min = 999999999999999999
	for i := 0; i < len(positions); i++ {
		var currentMin int
		for _, v := range positions {
			result := int(math.Abs(float64(v - i)))
			tmp := (result * (result + 1)) / 2
			currentMin += tmp
		}
		if currentMin < min {
			min = currentMin
		}
	}
	fmt.Println("min", min)
	if err != nil {
		fmt.Println("error", err)
	}
	// fmt.Println("max", max)
	// part1(card, numbers)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
