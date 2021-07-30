package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var myArray []int
	for scanner.Scan() {
		currentValue, _ := strconv.Atoi(scanner.Text())
		myArray = append(myArray, currentValue)
		fmt.Println(scanner.Text())
	}
	sort.Ints(myArray)
	if err != nil {
		fmt.Println("error", err)
	}
	for i := 0; i < len(myArray); i++ {
		for j := i + 1; j < len(myArray); j++ {
			for x := j + 1; x < len(myArray); x++ {
				if myArray[i]+myArray[j]+myArray[x] == 2020 {
					fmt.Println("first", myArray[i], "second", myArray[j], "third", myArray[x])
				}
			}
		}
	}
	fmt.Println(myArray)
}
