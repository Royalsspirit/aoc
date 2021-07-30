package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	countTrue := 0
	var myArray [][]string
	for scanner.Scan() {
		currentValue := scanner.Text()
		myArray = append(myArray, strings.Split(currentValue, ""))
		fmt.Println(strings.Split(currentValue, ""))
	}
	var slopes [][]int = [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}
	fmt.Println("len", len(myArray))
	for _, v := range slopes {

		var previousCellPos int = v[0]
		var localConut int = 0

		var positionSaved []string
		for i := 0; i < len(myArray)-v[1]; {

			positionSaved = append(positionSaved, "y="+strconv.Itoa(i+v[1])+", x="+strconv.Itoa(previousCellPos))
			// try to optimise this by removing the string addition
			// try to start from beginning when the tail is met
			if len(myArray[i+v[1]])-1 < previousCellPos {
				tmp := strings.Join(myArray[i+v[1]], "")
				nbrToRepeat := previousCellPos - (len(myArray[i+v[1]]) - 1)
				repeter := strings.Repeat(tmp, nbrToRepeat)
				sliced := strings.Split(repeter, "")
				myArray[i+v[1]] = append(myArray[i+v[1]], sliced...)
			}
			if myArray[i+v[1]][previousCellPos] == "#" {
				localConut++
				countTrue++
			}
			previousCellPos += v[0]
			// start from previous position
			i = i + v[1]

			// fmt.Println("position", positionSaved, "len", len(positionSaved))
		}
		fmt.Println("tree found ", localConut, "in current slopes", v)
	}
	fmt.Println("total", countTrue)
}
