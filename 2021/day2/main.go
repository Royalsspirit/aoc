package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(myArray map[string]int) int {
	return (myArray["down"] - myArray["up"]) * myArray["forward"]
}

func part2(myArray map[string]int) int {
	return myArray["forward"] * myArray["depth"]
}

func main() {
	f, err := os.Open("input-ex.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	myArray := make(map[string]int)
	for scanner.Scan() {
		currentValue := scanner.Text()
		move := strings.Split(currentValue, " ")
		unit, _ := strconv.Atoi(move[1])
		fmt.Println("move", move, " unit", unit)
		if _, found := myArray["aim"]; found && move[0] == "forward" {
			myArray["depth"] += myArray["aim"] * unit
		}

		myArray[move[0]] += unit

		fmt.Println("myarray", myArray)
		if move[0] == "up" {
			myArray["aim"] -= unit
		}
		if move[0] == "down" {
			myArray["aim"] += unit
		}

	}
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(part1(myArray))
	fmt.Println(part2(myArray))
}
