package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(myArray []int) int {
	increase := 0
	for i := 1; i < len(myArray); i++ {
		if myArray[i-1] < myArray[i] {
			increase = increase + 1
		}

	}

	return increase

}

/**
* sum the three next number and compare it to the next first number
* A: 607 (N/A - no previous sum)
B: 618 (increased)
C: 618 (no change)
D: 617 (decreased)
E: 647 (increased)
F: 716 (increased)
G: 769 (increased)
H: 792 (increased)
**/
func subSum(myArray []int, index int) int {
	return myArray[index] + myArray[index+1] + myArray[index+2]
}
func part2(myArray []int) int {
	increase := 0
	for i := 1; i < len(myArray)-2; i++ {
		fmt.Println("i", myArray[i])
		if subSum(myArray, i-1) < subSum(myArray, i) {
			increase = increase + 1
		}
	}
	return increase
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var myArray []int
	for scanner.Scan() {
		currentValue, _ := strconv.Atoi(scanner.Text())
		myArray = append(myArray, currentValue)
	}
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(part1(myArray))
	fmt.Println(part2(myArray))
}
