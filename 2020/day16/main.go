package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rules struct {
	name string
	min  int
	max  int
	min2 int
	max2 int
}

func part1(tickets [][]string, notes []rules) (int, map[int]int) {
	var errorsValues int = 0
	indexesFailed := make(map[int]int)
	for k, _ := range tickets {
		for _, vv := range tickets[k] {
			currentIntValue, _ := strconv.Atoi(vv)
			var matchRules bool = false
			for _, rule := range notes {
				if (rule.min <= currentIntValue && rule.max >= currentIntValue) ||
					(rule.min2 <= currentIntValue && rule.max2 >= currentIntValue) {
					matchRules = true
					break
				}

			}
			if !matchRules {
				indexesFailed[k] = k
				errorsValues += currentIntValue
			}
		}
	}
	return errorsValues, indexesFailed

}
func part2(tickets [][]string, notes map[string][]rules, failure map[int]int) map[string]map[string]int {
	matches := make(map[string]map[int]int)
	for k, _ := range tickets {
		if _, found := failure[k]; !found {
			for postTicket, vv := range tickets[k] {
				currentIntValue, _ := strconv.Atoi(vv)
				for posNote, rule := range notes {
					if rule[0].min <= currentIntValue && rule[0].max >= currentIntValue {
						if _, found := matches[posNote]; found {
							matches[posNote][postTicket]++
						} else {
							matches[posNote] = make(map[int]int)
							matches[posNote][postTicket] = 1
						}
					}
					if rule[1].min <= currentIntValue && rule[1].max >= currentIntValue {
						if _, found := matches[posNote]; found {
							matches[posNote][postTicket]++
						} else {
							matches[posNote] = make(map[int]int)
							matches[posNote][postTicket] = 1
						}
					}

				}
				fmt.Println("matches", matches, "for current value", currentIntValue)

			}
		}

	}
	fmt.Println("matches", matches)
	result := make(map[string]map[string]int)
	for k, _ := range matches {
		max := make(map[string]int)
		for posMax, vv := range matches[k] {
			if max["value"] < vv {
				max["value"] = vv
				max["position"] = posMax
			}
		}
		result[k] = max
	}
	return result
}

func part22(tickets [][]string, notes []rules, failure map[int]int) map[string]int {

	matches := make(map[string][]int)
	result := make(map[string]int)
	orderByColumn := make(map[int][]int)
	saved := make(map[int]int)
	for k, _ := range tickets {
		for kk, vv := range tickets[k] {
			if _, found := failure[k]; !found {
				ticketValueInt, _ := strconv.Atoi(vv)
				orderByColumn[kk] = append(orderByColumn[kk], ticketValueInt)
			}
		}
	}
	fmt.Println("target", len(orderByColumn[0]))
	for _, rule := range notes {
		i := 0
		for len(matches[rule.name]) != 190 {
			matches[rule.name] = []int{}
			if _, found := saved[i]; !found {
				for _, v := range orderByColumn[i] {
					currentIntValue := v
					if (rule.min <= currentIntValue && rule.max >= currentIntValue) ||
						(rule.min2 <= currentIntValue && rule.max2 >= currentIntValue) {
						matches[rule.name] = append(matches[rule.name], v)
					}
				}

			}
			fmt.Println("position i ", i, "len matche is", len(matches[rule.name]))
			if len(matches[rule.name]) == 190 {
				fmt.Println("i", i, " is position for", rule.name)
			}
			if i > len(orderByColumn[i]) {
				fmt.Println("len orderbycolumn when break", len(orderByColumn))
				break
			}
			i++
		}
		if len(matches[rule.name]) == 190 {
			saved[i-1] = i - 1
			result[rule.name] = i - 1

		} else {
			fmt.Println("add")
			notes = append(notes, rule)
		}
	}
	fmt.Println("result", result)
	return result

}

/**
* need to be update
* algorithm used:
* store every column position where each value is validated by the current rule
* for loop over the previous result and loop up for rules with single value and remove it to other rules
* at this end, we obtain a map with a single position by column
**/
func part23(tickets [][]string, notes []rules, failure map[int]int) map[string]int {
	matches := make(map[string][]int)
	result := make(map[string]int)
	orderByColumn := make(map[int][]int)
	//saved := make(map[int]int)
	for k, _ := range tickets {
		for kk, vv := range tickets[k] {
			if _, found := failure[k]; !found {
				ticketValueInt, _ := strconv.Atoi(vv)
				orderByColumn[kk] = append(orderByColumn[kk], ticketValueInt)
			}
		}
	}
	for i := 0; i < 20; i++ {
		for _, rule := range notes {
			var matchAll = true
			for _, v := range orderByColumn[i] {
				currentIntValue := v
				if !(rule.min <= currentIntValue && rule.max >= currentIntValue) &&
					!(rule.min2 <= currentIntValue && rule.max2 >= currentIntValue) {
					matchAll = false
				}
			}
			if matchAll {
				matches[rule.name] = append(matches[rule.name], i)
			}

		}
	}
	var rulesValue []string
	for ruleV, _ := range matches {
		rulesValue = append(rulesValue, ruleV)
	}
	// find other way to filter value
	for j := 0; j < len(rulesValue); j++ {
		if len(matches[rulesValue[j]]) == 1 {
			value := matches[rulesValue[j]][0]
			for i := 0; i < len(rulesValue); i++ {
				currentRuleValue := matches[rulesValue[i]]
				ruleName := rulesValue[i]
				if len(currentRuleValue) > 1 {
					for kkk, vvv := range currentRuleValue {
						if vvv == value && rulesValue[j] != rulesValue[i] {
							matches[ruleName] = append(matches[ruleName][:kkk], matches[ruleName][kkk+1:]...)
							j = 0
							break
						}
					}
				}

			}
		}
	}
	for k, v := range matches {
		result[k] = v[0]
	}
	return result

}

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	notes := []rules{}
	var myTicketValues []string
	var nearTicketsToTest [][]string
	var filterRules = true
	var nearbyTicket = false
	var myTicket = false
	for scanner.Scan() {

		currentValue := scanner.Text()
		if currentValue == "" {
			filterRules = false
		}
		if currentValue == "nearby tickets:" {
			nearbyTicket = true
			continue
		}
		if currentValue == "your ticket:" {
			myTicket = true
			continue
		}
		if myTicket {
			tickets := strings.Split(currentValue, ",")
			myTicketValues = append(myTicketValues, tickets...)
			myTicket = false

		}
		if filterRules {
			note := strings.Split(currentValue, ": ")
			currentRule := strings.Split(note[1], " or ")
			firstRule := strings.Split(currentRule[0], "-")
			secondRules := strings.Split(currentRule[1], "-")
			var r rules
			r.min, _ = strconv.Atoi(firstRule[0])
			r.max, _ = strconv.Atoi(firstRule[1])
			r.min2, _ = strconv.Atoi(secondRules[0])
			r.max2, _ = strconv.Atoi(secondRules[1])
			r.name = note[0]
			notes = append(notes, r)

		}
		if nearbyTicket {
			tickets := strings.Split(currentValue, ",")
			nearTicketsToTest = append(nearTicketsToTest, tickets)
		}
	}
	errorCount, failureTciket := part1(nearTicketsToTest, notes)
	fmt.Println("failure ticket", errorCount)
	matches := part23(nearTicketsToTest, notes, failureTciket)
	var finalValue int = 1
	for k, v := range matches {
		if strings.Index(k, "departure") != -1 {
			fmt.Println("k", k, "at position", v, "value is", myTicketValues[v])
			valueTicketInt, _ := strconv.Atoi(myTicketValues[v])
			finalValue *= valueTicketInt
		}
	}
	fmt.Println("final response part 2", finalValue)
}
