package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	var myArray []map[string]string
	tmpArray := make(map[string]string)
	var count int = 0
	for scanner.Scan() {
		currentValue := scanner.Text()
		if currentValue == "" {
			myArray = append(myArray, tmpArray)
			tmpArray = make(map[string]string)
			continue
		}
		passport := strings.Split(currentValue, " ")
		// find a better way to explit current value
		for _, v := range passport {
			values := strings.Split(v, ":")
			tmpArray[values[0]] = values[1]
		}

	}
	var mandatory [][]string = [][]string{
		[]string{"byr", "19[2-9][0-9]|200[0-2]"},
		[]string{"iyr", "201[0-9]|2020"},
		[]string{"eyr", "202[0-9]|2030"},
		[]string{"hgt", "^59in|[6-7][0-9]in$|^1[5-8][0-9]cm|19[0-3]cm$"},
		[]string{"hcl", "#[0-9a-f]{6}"},
		[]string{"ecl", "amb|blu|brn|gry|grn|hzl|oth"},
		[]string{"pid", "^\\d{9}$"},
	}
	// add last item
	myArray = append(myArray, tmpArray)
	for _, v := range myArray {
		for _, vv := range mandatory {
			matched, _ := regexp.MatchString(vv[1], v[vv[0]])
			if v[vv[0]] == "" || !matched {
				count++
				fmt.Println("not found ", vv, "in", v)
				break
			}
		}
	}
	fmt.Println("batch", len(myArray))
	fmt.Println("count", len(myArray)-count)
}
