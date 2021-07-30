package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	var myArray [][]string

	var tmpArray []string
	var result []int
	var sum int = 0
	for scanner.Scan() {
		currentValue := scanner.Text()
		if len(currentValue) == 0 {
			myArray = append(myArray, tmpArray)
			tmpArray = []string{}
		} else {
			tmpArray = append(tmpArray, currentValue)
		}
	}

	myArray = append(myArray, tmpArray)
	for _, v := range myArray {
		tmp := make(map[string]int)
		var max int = 0
		for _, vv := range v {
			if len(v) > 1 {
				// loop over rune
				for _, vvv := range vv {
					tmp[string(vvv)]++
					if tmp[string(vvv)] == len(v) {
						max++
					}
				}

			} else {
				max = len(vv)
			}
		}
		result = append(result, max)
	}
	for _, v := range result {
		sum = sum + v
	}
	fmt.Println("result", sum)
}
