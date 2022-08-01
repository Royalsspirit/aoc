package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(fishes *[]int) {
	currentLen := len(*fishes)
	for i := 0; i < currentLen; i++ {
		if (*fishes)[i] == 0 {
			(*fishes)[i] = 6
			*fishes = append(*fishes, 8)
		} else {
			(*fishes)[i]--
		}
	}
	fmt.Println("localfishes", fishes)
}

func part2(fishes *map[int]int) {
	var newFishes int
	var reset int
	for i := 0; i < 9; i++ {
		fmt.Println("i", i, "fishes", fishes)
		if _, found := (*fishes)[i]; found {
			if i == 0 {
				reset = (*fishes)[0]
				newFishes = (*fishes)[0]
			} else {
				(*fishes)[i-1] += (*fishes)[i]
			}
			delete((*fishes), i)

		}

	}

	(*fishes)[6] += reset
	(*fishes)[8] += newFishes
	tmpSum := 0
	for _, v := range *fishes {
		tmpSum += v
	}
	fmt.Println("tmpsum", tmpSum)
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var fishes []int
	fishes2 := make(map[int]int)
	for scanner.Scan() {

		currentValue := scanner.Text()
		currentFishes := strings.Split(currentValue, ",")
		for _, v := range currentFishes {
			vInt, _ := strconv.Atoi(v)
			fishes = append(fishes, vInt)
			fishes2[vInt]++
		}
	}

	for i := 0; i < 80; i++ {
		part1(&fishes)
	}
	fmt.Println("part1len", len(fishes))

	fmt.Println("fishes2", fishes2)
	for i := 0; i < 256; i++ {
		fmt.Println("dayyyyyy", i)
		part2(&fishes2)
	}

	fmt.Println("fishes", fishes, "len", len(fishes))
	sum := 0
	for _, v := range fishes2 {
		sum += v
	}
	fmt.Println("sum", sum)
	if err != nil {
		fmt.Println("error", err)
	}
	// fmt.Println("max", max)
	// part1(card, numbers)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
