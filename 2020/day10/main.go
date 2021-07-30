package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func tribonnaci(init int, limit int, adapter map[int]int) int {
	res := 0
	if init == limit {
		return 1
	}
	if _, found := adapter[init+1]; found {
		res += tribonnaci(init+1, limit, adapter)
	}
	if _, found := adapter[init+2]; found {
		res += tribonnaci(init+2, limit, adapter)
	}
	if _, found := adapter[init+3]; found {
		res += tribonnaci(init+3, limit, adapter)
	}
	return res
}

// solution 259172170858496
// work with small numbers
func part2(diffOnes int) int {
	total := diffOnes
	fmt.Println("total", total)
	i := 0
	for i < diffOnes {
		total += diffOnes - i
		i++
	}
	return total
}

func part22(values map[int]int, max int) map[int]int {
	index := 0
	path := make(map[int]int)
	path[0] = 1
	for index != max {
		fmt.Println("current digit", values[index])
		if _, found := values[index]; found || index == 0 {
			fmt.Println("found", index, "in", values)
			if _, found := values[index+1]; found {
				fmt.Println("found", values[index+1], "at index", index+1)
				fmt.Println("path is", path)
				path[values[index+1]] += path[values[index]]
			}
			if _, found := values[index+2]; found {

				fmt.Println("found", values[index+2], "at index", index+2)
				path[values[index+2]] += path[values[index]]
			}
			if _, found := values[index+3]; found {

				fmt.Println("found", values[index+3], "at index", index+3)
				path[values[index+3]] += path[values[index]]
			}
		}

		index = index + 1

	}
	return path
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	adapters := make(map[int]int)
	var maxAdapter int = 0
	for scanner.Scan() {
		currentValue := scanner.Text()
		valueInt, _ := strconv.Atoi(currentValue)
		adapters[valueInt] = valueInt
		if maxAdapter < valueInt {
			maxAdapter = valueInt
		}
	}
	adapters[maxAdapter+3] = maxAdapter + 3
	fmt.Println(adapters)
	start := 0
	jvolt := make(map[int]int)
	for start != maxAdapter+3 {
		firstA := start + 1
		second1 := start + 2
		thirdA := start + 3
		fmt.Println("start is", adapters[start])
		if _, found := adapters[firstA]; found {
			fmt.Println("here")
			jvolt[1]++
			start = firstA
		} else if _, found := adapters[second1]; found {
			jvolt[2]++
			start = second1
		} else if _, found := adapters[thirdA]; found {
			jvolt[3]++
			start = thirdA
		}
		fmt.Println("next start is", start)
	}
	fmt.Println("jvolt", jvolt)
	fmt.Println("sum", jvolt[1]*jvolt[3])
	//	fmt.Println("part 2", tribonnaci(0, maxAdapter+3, adapters))
	// fmt.Println("part", part2(jvolt[1]))
	fmt.Println("part 22", part22(adapters, maxAdapter+3))
}
