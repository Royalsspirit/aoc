package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
NNCB
NN NC CB
hh -> c
hh: 2 = HH HH
hc: 2 ch: 2 = HCH HCH = HC HC CH CH = HC: 2 CH: 2


recreate the pair array and keep the char counter up to date
**/
func insertPolymere(pairCount map[string]int, pairs map[string]string, count map[string]int) (map[string]int, map[string]int) {
	newPairCount := make(map[string]int)
	for k, v := range pairCount {
		if _, found := pairs[k]; found {
			count[pairs[k]] += v
			newPairCount[string(k[0])+pairs[k]] += v
			newPairCount[pairs[k]+string(k[1])] += v
		}
	}
	return newPairCount, count

}
func insert(template string, pairs map[string]string, count map[string]int) string {
	var p string = ""
	var c string
	for _, v := range template {
		s := string(v)
		p += s
		if _, found := pairs[p]; found && len(p) == 2 {
			if len(c) == 0 {

				c += string(p[0]) + pairs[p] + string(p[1])
				count[string(p[0])]++
				count[pairs[p]]++
				count[string(p[1])]++

			} else {

				c += pairs[p] + string(p[1])
				count[pairs[p]]++
				count[string(p[1])]++

			}
			p = string(p[1])
		}
	}
	return c

}
func part1(template string, pairs map[string]string) {
	var c string = template
	count := make(map[string]int)
	for i := 0; i < 40; i++ {
		count = make(map[string]int)
		c = insert(c, pairs, count)
	}
	var min, max int = 99999, 0
	for _, v := range count {
		if min > v {
			min = v
		}

		if v > max {
			max = v
		}
	}
	fmt.Println("max=", max, "min=", min, "equal", max-min)
}

/**
* instead of tracking the string, count the number of pair available
**/
func part2(template string, pairs map[string]string) {
	pairCount := make(map[string]int)
	count := make(map[string]int)
	for i := 0; i < len(template); i++ {
		count[string(template[i])]++
	}
	for i := 0; i < len(template)-1; i++ {
		pairCount[string(template[i])+string(template[i+1])]++
	}

	for i := 0; i < 40; i++ {
		pairCount, count = insertPolymere(pairCount, pairs, count)
	}

	var min, max int = 999999999999999999, 0
	for _, v := range count {
		if min > v {
			min = v
		}

		if v > max {
			max = v
		}
	}
	fmt.Println("max=", max, "min=", min, "equal", max-min)

}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	var template string
	pairs := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		currentValue := scanner.Text()
		if currentValue != "" {
			if strings.Index(currentValue, "->") == -1 {
				template = currentValue
			} else {
				pair := strings.Split(currentValue, " -> ")
				pairs[pair[0]] = pair[1]
			}
		}

	}
	if err != nil {
		fmt.Println("error", err)
	}
	// part1(template, pairs)
	part2(template, pairs)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
