package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func part1(chunks []string) {
	ass := map[string]string{
		"{": "}",
		"[": "]",
		"<": ">",
		"(": ")",
	}
	scorePart1 := map[string]int{
		"}": 1197,
		")": 3,
		"]": 57,
		">": 25137,
	}
	scorePart2 := map[string]int{
		"}": 3,
		")": 1,
		"]": 2,
		">": 4,
	}
	var chunksError []string
	var reminingChunk [][]string
	for _, v := range chunks {
		var stack []string
		chunksAsArray := strings.Split(v, "")
		for k, c := range chunksAsArray {
			fmt.Println("index", k)
			fmt.Println("stack", stack)
			_, found := ass[c]
			if len(stack) == 0 || (len(stack) > 0 && found) {
				stack = append([]string{ass[c]}, stack...)
			} else if c == stack[0] {
				stack = stack[1:]
			} else {
				fmt.Println("unexpected", c)
				chunksError = append(chunksError, c)
				stack = []string{}
				break
			}
		}
		if len(stack) != 0 {

			reminingChunk = append(reminingChunk, stack)
		}
	}
	finalScore := 0
	for _, v := range chunksError {
		finalScore += scorePart1[v]
	}
	fmt.Println("chunksError", finalScore)
	fmt.Println("remining chunk", reminingChunk)
	var finalScorePart2 []int
	for _, v := range reminingChunk {
		var tmp int
		for _, c := range v {
			tmp = (tmp * 5) + scorePart2[c]
		}
		fmt.Println("tmp", tmp)
		finalScorePart2 = append(finalScorePart2, tmp)
	}
	sort.Ints(finalScorePart2)
	fmt.Println("finalscorepart", finalScorePart2)
	fmt.Println("finalscorepart2", finalScorePart2[len(finalScorePart2)/2])
}

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	var chunks []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		currentValue := scanner.Text()
		chunks = append(chunks, currentValue)

	}
	if err != nil {
		fmt.Println("error", err)
	}
	part1(chunks)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
