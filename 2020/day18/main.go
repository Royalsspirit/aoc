package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://en.wikipedia.org/wiki/Reverse_Polish_notation (no order)
// should implement shuting yard algorithm. See part 2
func reversePolishNotation(queueOfOperators []string, queueOfOperands []int, tabString []string) (int, []string, []int) {

	for _, vv := range tabString {
		savedValue := vv
		interVal, err := strconv.Atoi(vv)
		if err != nil && (savedValue == "*" || savedValue == "+") {
			queueOfOperators = append(queueOfOperators, savedValue)
			continue
		}
		queueOfOperands = append(queueOfOperands, interVal)

		if len(queueOfOperands) == 2 {
			a := queueOfOperands[0]
			queueOfOperands = queueOfOperands[1:]
			b := queueOfOperands[0]
			queueOfOperands = queueOfOperands[1:]
			op := queueOfOperators[0]
			queueOfOperators = queueOfOperators[1:]
			if op == "+" {
				queueOfOperands = append(queueOfOperands, (a + b))
			} else if op == "*" {
				queueOfOperands = append(queueOfOperands, (a * b))
			}
		}
	}

	return 0, queueOfOperators, queueOfOperands
}

// Shunting-yard (with order)
// addition has higher precedance. Focus on it
func part2(queueOfOperators []string, queueOfOperands []int, tabString []string) (int, []string, []int) {

	var waitForOperands bool = true
	fmt.Println("current tab string", tabString)
	for _, vv := range tabString {
		savedValue := vv
		if vv == "+" {
			waitForOperands = false
		}
		interVal, err := strconv.Atoi(vv)
		if err != nil && (savedValue == "*" || savedValue == "+") {
			queueOfOperators = append(queueOfOperators, savedValue)
			continue
		}

		queueOfOperands = append(queueOfOperands, interVal)
		if len(queueOfOperands) >= 2 && !waitForOperands {
			fmt.Println("queueOfOperands", queueOfOperands)
			a := queueOfOperands[len(queueOfOperands)-1]

			queueOfOperands = queueOfOperands[:len(queueOfOperands)-1]

			fmt.Println("after dequeue queueOfOperands", queueOfOperands)
			b := queueOfOperands[len(queueOfOperands)-1]
			queueOfOperands = queueOfOperands[:len(queueOfOperands)-1]

			queueOfOperators = queueOfOperators[:len(queueOfOperators)-1]
			queueOfOperands = append(queueOfOperands, (a + b))
			waitForOperands = true
		}
	}
	fmt.Println("before --- queueOfOperators", queueOfOperators, "queueOfOperands", queueOfOperands)
	for len(queueOfOperators) > 0 {
		a := queueOfOperands[len(queueOfOperands)-1]
		queueOfOperands = queueOfOperands[:len(queueOfOperands)-1]
		b := queueOfOperands[len(queueOfOperands)-1]
		queueOfOperands = queueOfOperands[:len(queueOfOperands)-1]

		queueOfOperators = queueOfOperators[:len(queueOfOperators)-1]
		queueOfOperands = append(queueOfOperands, (a * b))

	}

	fmt.Println("after --- queueOfOperators", queueOfOperators, "queueOfOperands", queueOfOperands)
	return 0, queueOfOperators, queueOfOperands
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	var numbers []string
	for scanner.Scan() {
		currentValue := scanner.Text()
		numbers = append(numbers, currentValue)

	}
	r, _ := regexp.Compile(`\([^\(\)]+\)`)

	for i := 0; i < len(numbers); i++ {
		match := r.MatchString(numbers[i])
		if match {
			rawResult := r.FindAll([]byte(numbers[i]), -1)
			var queueOfOperators []string
			var queueOfOperands []int

			for _, b := range rawResult {
				v := string(b)
				// remove bracket
				v = v[1 : len(v)-1]
				tabOfResult := strings.Fields(v)
				if len(tabOfResult) > 1 {
					_, queueOfOperators, queueOfOperands = part2(queueOfOperators, queueOfOperands, tabOfResult)
					acc := strconv.Itoa(queueOfOperands[0])
					// fucking ligne 115 of input which have two same operation in same line
					numbers[i] = strings.Replace(numbers[i], string(b), acc, -1)
					queueOfOperands = queueOfOperands[1:]

				} else {
					fmt.Println("found one element", string(b), "in", numbers[i])
					numbers[i] = strings.Replace(numbers[i], string(b), v, -1)
				}

			}
			// restart for beginning
			i = -1
		}

	}
	fmt.Println("numbers", numbers)
	var final int = 0
	for j := 0; j < len(numbers); j++ {
		currentOperation := strings.Fields(string(numbers[j]))
		var queueOfOperands []int
		var queueOfOperators []string
		_, queueOfOperators, queueOfOperands = part2(queueOfOperators, queueOfOperands, currentOperation)
		final += queueOfOperands[0]
	}
	fmt.Println("part 2", final)
	//fmt.Println("final response part 1", numbersWithoutBraket)
}
