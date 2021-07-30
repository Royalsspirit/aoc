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
	for scanner.Scan() {
		currentValue := scanner.Text()
		arrayPassword := strings.Split(currentValue, ": ")
		rules := strings.Split(arrayPassword[0], " ")
		count := strings.Split(rules[0], "-")
		min, _ := strconv.Atoi(count[0])
		max, _ := strconv.Atoi(count[1])
		min = min - 1
		max = max - 1
		letter := rules[1]
		password := arrayPassword[1]

		/*
			part 1
			pattern := regexp.MustCompile(letter)
			matched := pattern.FindAllStringIndex(password, -1)
			if len(matched) >= min && len(matched) <= max {
				countTrue++
			}
			fmt.Println(matched, "current password", password, "rules", rules)
		*/

		/* part 2 */
		if (string(password[min]) == letter && string(password[max]) != letter) || string(password[max]) == letter && string(password[min]) != letter {
			countTrue++
		}
	}
	fmt.Println("coutn true", countTrue)
}
