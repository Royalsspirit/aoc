package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findSumMember(preamble []int, result int) bool {
	for _, first := range preamble {
		for _, second := range preamble {
			if first+second == result {
				return true
			}
		}
	}
	return false

}

func vulnerability(values []int, badNumber int) int {
	min, max := 0, 1
	intervalSum := values[min] + values[max]
	for intervalSum != badNumber {
		fmt.Println(intervalSum)
		if intervalSum < badNumber {
			max++
			intervalSum += values[max]
		}
		if intervalSum > badNumber {
			intervalSum -= values[min]
			min++
		}
	}
	contigousRange := values[min : max+1]

	minRange := contigousRange[0]
	maxRange := 0
	for _, v := range contigousRange {
		fmt.Println("v", v)
		if minRange > v {
			minRange = v
		}
		if maxRange < v {
			maxRange = v
		}
	}

	return minRange + maxRange
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	var values []int
	for scanner.Scan() {
		currentValue := scanner.Text()
		valueInt, _ := strconv.Atoi(currentValue)
		values = append(values, valueInt)
	}

	var badNumber int
	for i := 0; i < len(values); i++ {
		sumResult := values[i+25]
		if !findSumMember(values[i:25+i], sumResult) {
			badNumber = sumResult
			break
		}
	}
	sum := vulnerability(values, badNumber)
	fmt.Println("sum", sum)
}
